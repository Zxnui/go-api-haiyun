package controllers

import (
	"github.com/martini-contrib/render"
	"net/http"
)

func Home(r render.Render, rq *http.Request) {
	r.HTML(200, "home", nil)
}
