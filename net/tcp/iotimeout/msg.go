package iotimeout

import (
	"encoding/json"
	"fmt"
)

type Msg struct {
	Data string `json:"Data"`
}

func NewMsg(src, dst, data string) *Msg {
	return &Msg{
		Data: data,
	}
}

func Encode(m *Msg) []byte {
	bs, _ := json.Marshal(m)
	return bs
}

func Decode(bs []byte) *Msg {
	msg := &Msg{}
	_ = json.Unmarshal(bs, msg)
	return msg
}

func (m *Msg) Print() {
	fmt.Println(fmt.Sprintf("content:%s", m.Data))
}
