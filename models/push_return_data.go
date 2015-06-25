package models

import (
// "github.com/go-xorm/xorm"
)

type PushReturn struct {
	Pushdata   `xorm:"extends"`
	Returndata `xorm:"extends"`
}

func GetAllPushReturn() []PushReturn {
	ads := make([]PushReturn, 0)
	Xorm.Table("pushdata").Join("INNER", "returndata", "pushdata.grtzone_id = returndata.grtzone.id").Find(&ads)
	return ads
}
