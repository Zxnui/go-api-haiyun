package controllers

import (
	"github.com/martini-contrib/render"
	"go-api-haiyun/models"
	"net/http"
)

func Home(r render.Render, rq *http.Request) {
	data := make(map[string]interface{})
	data["AdsData"] = models.GetAllPushReturn()
	r.HTML(200, "home", data)
}
