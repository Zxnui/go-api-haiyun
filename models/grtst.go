package models

import (
	// "github.com/go-xorm/xorm"
	"time"
)

type Grtst struct {
	Id       int64
	Username string
	Passwd   string
	Role     int
	Ratio    float64
	SubDate  time.Time
	ReMark   string
}
