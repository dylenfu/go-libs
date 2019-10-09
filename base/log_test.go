package base

import (
	"log"
	"os"
	"testing"
)

var whitelist *Whitelist

// Whitelist .
type Whitelist struct {
	log  *log.Logger
	list map[int64]struct{} // whitelist for debug
}

func TestSimpleLog(t *testing.T) {
	path := "/tmp/white_list.log"
	list := []int64{12, 345}

	if f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644); err == nil {
		whitelist = new(Whitelist)
		whitelist.log = log.New(f, "", log.LstdFlags)
		whitelist.list = make(map[int64]struct{})
		for _, mid := range list {
			whitelist.list[mid] = struct{}{}
		}
	}

	t.Log(1)
}
