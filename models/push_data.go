package models

import (
	// "github.com/go-xorm/xorm"
	"time"
)

type Pushdata struct {
	Id        int64
	GrtzoneId int64
	Pname     string
	Weburl    string
	Puburl    string
	Pubtype   int
	Layout    int
	EffectId  int
	Width     int
	Height    int
	SubDate   time.Time
	EditTime  time.Time
}
