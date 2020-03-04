package sense

import (
	"encoding/json"
	"testing"
)

type SignTx struct {
	Data string `json:"data"`
}

type SignatureTxResp struct {
	Code       uint32  `json:"code"`     //响应代码(ErrorCode)
	SignedData *SignTx `json:"SignData"` //已经签好名的交易信息
}

// go test -v github.com/dylenfu/go-libs/sense -run TestUnmarshal
func TestUnmarshal(t *testing.T) {
	resp := &SignatureTxResp{SignedData: new(SignTx)}
	data := `{"code":12}`

	if err := json.Unmarshal([]byte(data), resp); err != nil {
		t.Log("-----failed")
	} else {
		t.Log("-----success")
		t.Log(resp.SignedData.Data)
	}
}
