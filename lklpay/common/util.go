package common

import (
	"time"
)

func GetReqTime() string {
	return time.Now().Format("20060102150405")
}
