// DEPRECATED BEACAUSE OF code.google.com/p/go-netrc/netrc 
package octokit

// import (
// 	"sync/atomic"
// 	"time"
// 	"errors"
// )

// type ConcurrencyStatus bool
// type ConcurrencyManager func(title, body, level string) error

// var (
// 	Done       *uint32 = new(uint32)
// 	BusyWorker *int32  = new(int32)
// )

// var (
//   octokit로거         = octokitLogger()
//   octokitLogger1     = octokitLogger()
//   octokitLogger2     = octokitLogger()
//   octokitLogger3     = octokitLogger()
//   )

// func octokitLogger() ConcurrencyManager {
// return func(title, body, level string) error {
//   atomic.AddInt32(BusyWorker, 1)
//   atomic.AddUint32(Done, 1)
//   go CreateOrUpdateIssue(title, body, level)
//   return nil
// }
// }
// // Concurrent.
// func Concurrent(err error) error {
// 	c := make(chan error)
// 	defer close(c)
// 	go func() {
// 		c <- err
// 	}()

// 	timeout := time.After(30 * time.Second)
// 	select {
// 	case err := <-c:
// 		atomic.AddInt32(BusyWorker, -1)
// 			return err
// 	case <-timeout:
// 	  err := errors.New("Log timed out.")
// 		return err
// 	}
// 	return errors.New("Invalid Request.")
// }

// // ConcurrencyAgent is loadbalancer of concurrencier.
// func ConcurrencyAgent(title, body, level string, replicas ...ConcurrencyManager) error {
// 	for {
// 		workingNow := atomic.LoadInt32(BusyWorker)
// 		if len(replicas) > int(workingNow) {
// 			break
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	c := make(chan error)

// 	concurrencyReplica := func(i int) {
// 		c <- replicas[i](title, body, level)
// 	}
// 	workingNow := atomic.LoadInt32(BusyWorker)
// 	go concurrencyReplica(int(workingNow))
// 	return <-c
// }

// // SendLog send logs to github repository.
// func SendLog(title, body, level string) error {
//   return Concurrent(ConcurrencyAgent(title, body, level,
//       octokit로거,
//       octokitLogger1,
//       octokitLogger2,
//       octokitLogger3))

// }
