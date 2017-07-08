package middlewares

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/core/models"
	"github.com/Yaoguais/gadmin/core/services"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if id := core.GetUserId(c); id == "" {
			c.Redirect(http.StatusFound, "/login")
			return nil
		}

		return next(c)
	}
}

func GetCurrentUser(c echo.Context) *models.AdminUserModel {
	id := core.GetUserId(c)
	if id != "" {
		if uid, err := strconv.ParseInt(id, 10, 64); err == nil {
			return services.AdminService.GetAdminUserByUid(uid)
		}
	}

	return nil
}
