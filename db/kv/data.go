package kv

import (
	"sync"
)

type UserData struct {
	Uid      int32     `json:"uid" redis:"uid"`
	Nickname string    `json:"nickname" redis:"nickname"`
	Gender   int32     `json:"gender" redis:"gender"`
	Verified int32     `json:"verified" redis:"verified"`
	Portrait string    `json:"portrait" redis:"portrait"`
	Exp      int32     `json:"exp" redis:"exp"`
	Level    int32     `json:"level" redis:"level"`
	Noble    NobleData `json:"noble" redis:"-"`
	Rank     int32     `json:"rank" redis:"-"`
	Score    int64     `json:"score" redis:"-"`
}

type NobleData struct {
	EndTime   int32 `json:"-" redis:"-"`
	EndTime64 int64 `json:"-" redis:"endTime"`
	RoomHide  int32 `json:"-" redis:"roomHide"`
	Level     int32 `json:"level" redis:"level"`
	Status    int32 `json:"status" redis:"status"`
	Weight    int32 `json:"weight" redis:"weight"`
}

type UserRoomData struct {
	rid  int32
	refs int32
	last int64
}

type User struct {
	mu    sync.RWMutex
	last  int64
	rooms []UserRoomData
	state int32 // 0: empty, 1: preparing, 2: ok, 3: down
	uid   int32
	ti    uint64 // timer index
	tidx  int    // timer slot index
	cd    int64
	data  UserData
	noble NobleData
}
