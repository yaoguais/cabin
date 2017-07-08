package demo

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/labstack/echo"
)

type DemoController struct {
}

func (*DemoController) GetIndex(c echo.Context) error {
	return core.Render(c, core.GetAppViewPath("demo/index.html"), nil)
}
