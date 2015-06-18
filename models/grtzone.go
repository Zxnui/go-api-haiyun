package models

import (
	// "github.com/go-xorm/xorm"
	"time"
)

type Grtzone struct {
	Id      int64
	Name    string
	AdsType int
	Size    string
	SubDate time.Time
}
