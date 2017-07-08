package main

import (
	appRoutes "github.com/Yaoguais/gadmin/app/routes"
	"github.com/Yaoguais/gadmin/core"
	coreRoutes "github.com/Yaoguais/gadmin/core/routes"
)

func main() {

	core.SetConfig(core.GetAppRoot() + "/config/config.toml")
	core.InitApp()
	coreRoutes.RegisterRoutes()
	appRoutes.RegisterRoutes()
	core.RunApp()
}
