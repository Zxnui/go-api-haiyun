package models

import (
	// "github.com/go-xorm/xorm"
	"encoding/json"
	"time"
)

type Returndata struct {
	Id        int64
	GrtzoneId int64  `json:"pubid"`
	Status    int    `json:"status"`
	Code      string `json:"code"`
	Reason    string `json:"reason"`
	SubDate   time.Time
	EditTime  time.Time
}

func GetDataByJson() {

}
