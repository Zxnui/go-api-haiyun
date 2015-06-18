package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	_ "go-api-haiyun/conf"
	_ "go-api-haiyun/models"
	"go-api-haiyun/routers"
)

func main() {
	m := martini.Classic()

	//Make sure to include the Gzip middleware above other middleware that alter the response body (like the render middleware).
	m.Use(gzip.All())

	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))

	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("mysession", store))

	routers.RouterInit(m)

	m.Run()
}
