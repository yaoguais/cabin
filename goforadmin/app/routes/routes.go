package routes

import (
	"github.com/Yaoguais/gadmin/app/controllers/demo"
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/core/middlewares"
)

func RegisterRoutes() {
	app := core.GetApp()
	demoController := demo.DemoController{}

	g := app.Group("", middlewares.Auth, middlewares.Access)
	g.GET("/demo", demoController.GetIndex)
}
