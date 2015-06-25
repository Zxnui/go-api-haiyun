package conf

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Cfg *goconfig.ConfigFile

//加载配置文件
func init() {
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件:%s", err)
	}
	Cfg = cfg
}
