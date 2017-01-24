package concurrency

import (
	"errors"
	"mime/multipart"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/dorajistyle/goyangi/util/config"
	"github.com/dorajistyle/goyangi/util/log"
)

type ConcurrencyStatus bool
type ConcurrencyManagerLegacy func(reader *multipart.Reader) ConcurrencyStatus

type ConcurrencyManager func(request *http.Request) Result

var (
	Done       *uint32 = new(uint32)
	BusyWorker *int32  = new(int32)
)

// Result is the struct that contain http status code and error.
type Result struct {
	Code  int
	Error error
}

// Concurrent.
func Concurrent(request *http.Request, result Result) (int, error) {
	c := make(chan Result)
	defer close(c)
	go func() {
		c <- result
	}()

	timeout := time.After(config.UploadTimeout)
	select {
	case res := <-c:
		log.Debugf("End of Upload : %v", res)
		workingNow := atomic.AddInt32(BusyWorker, -1)
		log.Debugf("All files are Done. Working concurrencyer count : %d", workingNow)
		if workingNow == 0 {
			return res.Code, res.Error
		}
	case <-timeout:
		err := errors.New("Request timed out.")
		log.Warnf(err.Error())
		return http.StatusBadRequest, err
	}
	return http.StatusBadRequest, errors.New("Invalid Request.")
}

// ConcurrencyAgent is loadbalancer of concurrencier.
func ConcurrencyAgent(request *http.Request, replicas ...ConcurrencyManager) Result {
	for {
		workingNow := atomic.LoadInt32(BusyWorker)
		if len(replicas) > int(workingNow) {
			break
		}
		time.Sleep(time.Second)
		// log.Debugf("working concurrencyer count full (BusyWorker/replicas)  : (%d/%d) ", workingNow, len(replicas))
	}
	c := make(chan Result)

	concurrencyReplica := func(i int) {
		c <- replicas[i](request)
	}
	workingNow := atomic.LoadInt32(BusyWorker)
	log.Debugf("workingNow, len(replicas) : %d %d", workingNow, len(replicas))

	go concurrencyReplica(int(workingNow))
	// go concurrencyReplica(0)
	return <-c
}
