package models

import (
	// "github.com/go-xorm/xorm"
	"encoding/json"
	"log"
	"time"
)

type Pushdata struct {
	Id        int64
	GrtzoneId int64  `json:"pid"`
	Pname     string `json:"pname"`
	Weburl    string `json:"weburl"`
	Puburl    string `json:"puburl"`
	Pubtype   int    `json:"pubtype"`
	Layout    int    `json:"layout"`
	EffectId  int    `json:"effectId"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	SubDate   time.Time
	EditTime  time.Time
	Cpmprice  string `json:"cpmprice"`
	Cpcprice  string `json:"cpcprice"`
	Backjs    string `json:"backjs"`
}

//将数据转换为json格式
func GetJson(s Pushdata) (string, error) {
	json_data, err := simplejson.NewJson(s)
	if err != nil {
		return nil, err
	}
	return json_data, nil
}
