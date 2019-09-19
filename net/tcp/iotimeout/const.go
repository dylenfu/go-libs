package iotimeout

import "time"

const (
	tcp                         = "tcp"
	bufsize                     = 1024
	host                        = "127.0.0.1:9090"
	connname                    = "cli"
	writeDeadLine time.Duration = 200 * time.Millisecond // 100 ms
)
