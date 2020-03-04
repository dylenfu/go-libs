package sense

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

// 并发情况下，goroutine内进行加锁并不安全，取决于资源使用时间与调用时间
type service struct {
	metaLock      *sync.RWMutex
	withdrawBatch map[string]int
}

func (s *service) ApplyWithdrawal(batchID string, num int) (rsp string) {
	if !s.isWithDrawalBatch(batchID) {
		go s.DealWithApplyWithdrawal(batchID, num)
		rsp = fmt.Sprintf("batch id %s do not exist", batchID)
	} else {
		rsp = fmt.Sprintf("batch id %s exist already", batchID)
	}
	return
}

func (s *service) isWithDrawalBatch(batchID string) bool {
	s.metaLock.Lock()
	defer s.metaLock.Unlock()
	_, ok := s.withdrawBatch[batchID]
	return ok
}

func (s *service) DealWithApplyWithdrawal(batchID string, num int) {
	time.Sleep(1600 * time.Microsecond)
	s.metaLock.Lock()
	s.withdrawBatch[batchID] = num
	s.metaLock.Unlock()
}

// go test -v github.com/dylenfu/go-libs/sense -run TestUnsafeBatchIdWriting
func TestUnsafeBatchIdWriting(t *testing.T) {
	s := new(service)
	s.metaLock = new(sync.RWMutex)
	s.withdrawBatch = make(map[string]int)
	batchID := "test"

	for i := 0; i < 100; i++ {
		fmt.Println(s.ApplyWithdrawal(batchID, i))
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGSTOP)
	for {
		s := <-sig
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
