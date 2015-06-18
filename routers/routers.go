package routers

import (
	"github.com/go-martini/martini"
	"go-api-haiyun/controllers"
)

func RouterInit(m *martini.ClassicMartini) {
	m.Group("/home", func(r martini.Router) {
		r.Get("/index", controllers.Home)
	})
}
