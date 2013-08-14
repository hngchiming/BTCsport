package models

import (
	"time"
)

func Token() int64 {
	return time.Now().Add(time.Second * 10086).UnixNano()
}
