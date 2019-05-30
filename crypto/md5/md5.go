package md5

import(
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func Md5() {
	// data := []byte("hello world")
	// s := fmt.Sprintf("%x", md5.Sum(data))
	// fmt.Println(s)

	// 也可以用这种方式
	data := []byte("hello world")
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	fmt.Println(s)
}
