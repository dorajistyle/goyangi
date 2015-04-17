package email

import (
	"errors"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/util/log"
	gomail "gopkg.in/gomail.v1"
)

type EmailError error
type PostOfficer func(msg *gomail.Message) EmailError

var (
	mailer                 = InitGomail()
	sent           *uint32 = new(uint32)
	workingOfficer *int32  = new(int32)
	집배원                    = realPostOfficer()
	Postman                = realPostOfficer()
	Cartero                = realPostOfficer()
	Facteur                = realPostOfficer()
	Yóudìyuán              = realPostOfficer()
	почтальон              = realPostOfficer()
	Leterportisto          = realPostOfficer()
)

func InitGomail() *gomail.Mailer {
	mailer := gomail.NewMailer(config.EmailHost, config.EmailUsername, config.EmailPassword, config.EmailPort)
	return mailer
}

func SendEmail(msg *gomail.Message) (emailError EmailError) {
	c := make(chan EmailError)
	go func() {
		c <- PostOffice(msg, 집배원, Postman, Cartero, Facteur, Yóudìyuán, почтальон, Leterportisto)
	}()
	timeout := time.After(config.EmailTimeout)
	for i := 0; i < 3; i++ {
		select {
		case emailError = <-c:
			workingNow := atomic.AddInt32(workingOfficer, -1)
			log.Debugf("Email sent. Working officer count : %d", workingNow)
			// atomic.StoreUint32(sent, 0)
			return
		case <-timeout:
			// atomic.StoreUint32(sent, 0)
			workingNow := atomic.AddInt32(workingOfficer, -1)
			log.Debugf("Email sent. Working officer count : %d", workingNow)
			log.Warn("Email timed out")
			emailError = errors.New("Email timed out")
			return
		}
	}
	return
}

func PostOffice(msg *gomail.Message, postOfficers ...PostOfficer) EmailError {
	for {
		workingNow := atomic.LoadInt32(workingOfficer)
		if len(postOfficers) > int(workingNow) {
			break
		}
		time.Sleep(time.Second)
		// log.Debugf("working uploader count full (workingOfficer/replicas)  : (%d/%d) ", workingNow, len(replicas))
	}

	c := make(chan EmailError)
	postOfficer := func(i int) { c <- postOfficers[i](msg) }

	workingNow := atomic.LoadInt32(workingOfficer)
	log.Debugf("workingNow, len(replicas) : %d %d", workingNow, len(postOfficers))

	go postOfficer(int(workingNow))

	// for i := range postOfficers {
	// 	go postOfficer(i)
	// }
	return <-c
}

func realPostOfficer() PostOfficer {
	return func(msg *gomail.Message) EmailError {
		// if atomic.CompareAndSwapUint32(sent, 0, 1) {
		atomic.AddInt32(workingOfficer, 1)
		return mailer.Send(msg)
		// sentNow := atomic.LoadUint32(sent)
		// log.Debugf("sent now : %d", sentNow)
		// return nil
		// }
		return nil
	}
}

func SendEmailFromAdmin(to string, subject string, body string, bodyHTML string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.EmailFrom)
	msg.SetHeader("To", to, config.EmailTestTo)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)
	msg.AddAlternative("text/html", bodyHTML)
	log.Debugf("to : %s", to)
	log.Debugf("subject : %s", subject)
	log.Debugf("body : %s", body)
	log.Debugf("bodyHTML : %s", bodyHTML)
	if config.SendEmail {
		log.Debug("SendEmail performed.")
		err := SendEmail(msg)
		return err
	}
	return nil
}

func SendTestEmail() error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", config.EmailFrom)
	msg.SetHeader("To", config.EmailTestTo)
	msg.SetAddressHeader("Cc", config.EmailTestTo, "dorajistyle")
	msg.SetHeader("Subject", "Hi(안녕하세요)?!")
	msg.SetBody("text/plain", "Hi(안녕하세요)?!")
	msg.AddAlternative("text/html", "<p><b>Goyangi(고양이)</b> means <i>cat</i>!!?</p>")
	path, err := filepath.Abs("frontend/canjs/static/images/goyangi.jpg")
	if err != nil {
		panic(err)
	}
	f, err := gomail.OpenFile(path)
	if err != nil {
		panic(err)
	}
	msg.Attach(f)
	// SendEmail(msg)
	err = SendEmail(msg)
	return err
}
