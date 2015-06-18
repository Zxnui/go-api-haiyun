package models

import (
	// "github.com/go-xorm/xorm"
	"time"
)

type Returndata struct {
	Id       int64
	PubId    int
	Status   int
	Code     string
	Reason   string
	SubDate  time.Time
	EditTime time.Time
}
