package middlewares

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/labstack/echo"
)

func Admin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if u := GetCurrentUser(c); u != nil && u.Role == SuperUserRole {
			return next(c)
		}

		if c.Request().Header.Get("X-Requested-With") == "XMLHttpRequest" {
			return core.Response(c, &core.Resp{
				Ret: 1,
				Msg: "没有权限!",
			})
		} else {
			return core.ShowErrorPage(c, 403, "没有权限!")
		}

		return nil
	}
}
