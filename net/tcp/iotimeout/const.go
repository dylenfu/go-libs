package iotimeout

import "time"

const (
	tcp                         = "tcp"
	bufsize                     = 1024
	host                        = "127.0.0.1:9090"
	clientaddr                  = "127.0.0.1:9091"
	connname                    = "cli"
	writeDeadLint time.Duration = 100 * time.Millisecond // 100 ms
)
