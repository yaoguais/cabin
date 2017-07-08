package index

import (
	"github.com/Yaoguais/gadmin/lib/session"
	"github.com/labstack/echo"
	"net/http"
)

type LogoutController struct {
}

func (*LogoutController) GetLogout(c echo.Context) error {
	s := session.GetSession(c, session.SessionName())
	s.Options.MaxAge = -1
	session.SetSession(c, s)

	return c.Redirect(http.StatusFound, "/login")
}
