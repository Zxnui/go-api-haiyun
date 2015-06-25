package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"go-api-haiyun/conf"
	_ "go-api-haiyun/models"
	"go-api-haiyun/routers"
)

func main() {
	m := martini.Classic()

	//gzip必须放在所有中间件前面
	m.Use(gzip.All())
	m.Use(martini.Static("static"))

	//前端基本配置
	m.Use(render.Renderer(render.Options{
		Directory:  "views",                    //自定义前端模板文件
		Layout:     "layout",                   //设置母模板
		Extensions: []string{".tmpl", ".html"}, //模板格式
		Charset:    "UTF-8",                    //编码方式
	}))

	store := sessions.NewCookieStore([]byte("instreet"))
	m.Use(sessions.Sessions("instreet", store))

	routers.RouterInit(m)

	//读取端口，配置端口
	port, _ := conf.Cfg.GetValue("", "httpport")
	m.RunOnAddr(":" + port)
}
