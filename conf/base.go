package conf

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Cfg *goconfig.ConfigFile

func init() {
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件:%s", err)
	}
	Cfg = cfg
}
