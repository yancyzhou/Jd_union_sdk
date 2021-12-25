package pkg

import (
	"time"
)

const (
	TimeLayout string = "2006-01-02 15:04:05"
)

var (
	Location *time.Location
)

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}
