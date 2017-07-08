package index

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/labstack/echo"
)

type IndexController struct {
}

func (*IndexController) GetIndex(c echo.Context) error {
	return core.Render(c, core.GetCoreViewPath("index/index.html"), nil)
}
