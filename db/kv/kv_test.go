package kv

import (
	"testing"
)

func TestNewKVCachePool(t *testing.T) {
	p := NewKVCachePool(-1, 1, Shard)
	for i := 1; i < 100; i++ {
		uid := 150000 + i
		user := &UserData{
			Uid:      int32(uid),
			Nickname: "ddd",
			Gender:   0,
			Verified: 1,
			Portrait: "/ssssss",
			Exp:      1,
			Level:    1,
		}
		p.Set(uid, user, -1)
	}
	t.Log("finish")
}
