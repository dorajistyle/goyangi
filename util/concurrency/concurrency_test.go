package concurrency_test

import (
	"net/http"
	// "github.com/gin-gonic/gin"
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"sync/atomic"

	"github.com/dorajistyle/goyangi/util/concurrency"
	viper "github.com/dorajistyle/goyangi/util/viper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	viper.LoadConfig()
}

// mockCocurrencytManager
func mockCocurrencytManager() concurrency.ConcurrencyManager {
	return func(request *http.Request) concurrency.Result {
		atomic.AddInt32(concurrency.BusyWorker, 1)
		var result concurrency.Result
		result.Code = http.StatusOK
		reader, err := request.MultipartReader()

		// user, _ = userService.CurrentUser(c)

		if err != nil {
			result.Code = http.StatusInternalServerError
			result.Error = err
			return result
		}
		userAgent := request.Header.Get("User-Agent")
		var count int
		count = 0
		for {
			count += 1
			part, err := reader.NextPart()

			// uploadedNow := atomic.AddUint32(concurrency.Done, 1)

			if err == io.EOF || part == nil {

				break
			}
			if part.FormName() == "" {

				continue
			}
			result.Code, result.Error = ConcurrencyRun(part, userAgent)
		}
		return result
	}
}

func ConcurrencyRun(part *multipart.Part, userAgent string) (int, error) {
	var err error
	return 200, err
}

var _ = Describe("concurrency", func() {
	var (
		err          error
		req          *http.Request
		mockManager1 concurrency.ConcurrencyManager
		mockManager2 concurrency.ConcurrencyManager
		mockManager3 concurrency.ConcurrencyManager
		mockManager4 concurrency.ConcurrencyManager
		mockManager5 concurrency.ConcurrencyManager
	)

	BeforeEach(func() {
		mockManager1 = mockCocurrencytManager()
		mockManager2 = mockCocurrencytManager()
		mockManager3 = mockCocurrencytManager()
		mockManager4 = mockCocurrencytManager()
		mockManager5 = mockCocurrencytManager()
	})

	Describe("Run a concurrent", func() {

		Context("when concurrent finished", func() {
			BeforeEach(func() {

				req = &http.Request{
					Method: "POST",
					Header: http.Header{"Content-Type": {`multipart/form-data; boundary="foo123"`}},
					Body:   ioutil.NopCloser(new(bytes.Buffer)),
				}
				// multipart, err := req.MultipartReader()
				// if multipart == nil {
				//     t.Error("expected multipart;", err)
				// }

				// req.Header = Header{"Content-Type": {"text/plain"}}
				// multipart, err = req.MultipartReader()
				// if multipart != nil {
				//     t.Error("unexpected multipart for text/plain")
				// }

				_, err = concurrency.Concurrent(req, concurrency.ConcurrencyAgent(req, mockManager1, mockManager2, mockManager3, mockManager4, mockManager5))

			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})

		})
	})

})
