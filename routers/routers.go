package routers

import (
	"github.com/go-martini/martini"
	"go-api-haiyun/controllers"
)

func RouterInit(m *martini.ClassicMartini) {
	m.Get("/", controllers.Home)
}
