package index

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/core/services"
	"github.com/Yaoguais/gadmin/lib/session"
	tpl "github.com/Yaoguais/gadmin/lib/template"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type LoginController struct {
}

func (*LoginController) GetLogin(c echo.Context) error {
	id := core.GetUserId(c)
	if id != "" {
		return c.Redirect(http.StatusFound, "/")
	}

	d := tpl.Vars{
		"Title": "登录" + strings.ToUpper(core.GetAppName()) + "管理后台",
	}

	return core.Render(c, core.GetCoreViewPath("login/login.html"), d)
}

func (*LoginController) PostLogin(c echo.Context) error {
	var m *core.Resp

	username := c.FormValue("username")
	password := c.FormValue("password")

	m = &core.Resp{
		Ret: 1,
		Msg: "用户名或密码错误!",
	}

	if user := services.AdminService.GetAdminUserByUsername(username); user != nil {
		if services.AdminService.CheckPassword(user.Password, password) {
			// mark user as login
			s := session.GetSession(c, session.SessionName())
			services.AdminService.SaveUserSession(user, s)
			session.SetSession(c, s)

			m = &core.Resp{
				Ret: 0,
			}
		}
	}

	return core.Response(c, m)
}
