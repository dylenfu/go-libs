package base

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// go test -v github.com/dylenfu/go-libs/base -run TestBytesFormatString
func TestBytesFormatString(t *testing.T) {
	bs := []byte{'a', 'b', 'c', 'd'}
	for _, v := range bs {
		t.Log(v)
	}
	t.Log(fmt.Sprintf("%x", bs[:]))
	t.Log(fmt.Sprintf("discover.HexID(\"%x\")", bs[:]))
	t.Log(hex.EncodeToString(bs))
	t.Log(^uint16(0))
}
