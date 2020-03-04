package sense

import "testing"

/*
缓冲通道capacity有使用上限，当通道中的数据需要顺序处理时，如果处理速度跟不上输入速度，可能会导致丢失数据
虽然channel本身不能扩容，但是可以通过给通道数量进行扩容的方式处理并发请求.
*/
func TestChannelSharding(t *testing.T) {

}
