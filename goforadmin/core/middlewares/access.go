package middlewares

import (
	"fmt"
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/core/services"
	. "github.com/Yaoguais/gadmin/lib/string"
	"github.com/labstack/echo"
	"github.com/orcaman/concurrent-map"
	"strconv"
	"strings"
)

var (
	SuperUserId   int64 = 1
	SuperUserRole int64 = 1
	userRoleMap   cmap.ConcurrentMap
	userPrivMap   cmap.ConcurrentMap
)

func init() {
	userRoleMap = cmap.New()
	userPrivMap = cmap.New()
}

func Access(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if id := core.GetUserId(c); id != "" {
			uid, _ := strconv.ParseInt(id, 10, 64)
			if uid == SuperUserId {
				return next(c)
			}

			if role, ok := userRoleMap.Get(id); ok {
				if role.(int64) == SuperUserRole {
					return next(c)
				}
			}

			pk := fmt.Sprintf("%s %s %s", id, c.Request().Method, c.Request().RequestURI)
			pk = strings.ToLower(pk)

			if ret, ok := userPrivMap.Get(pk); ok && ret.(string) == "1" {
				return next(c)
			}
		}

		if c.Request().Header.Get("X-Requested-With") == "XMLHttpRequest" {
			core.Response(c, &core.Resp{
				Ret: 1,
				Msg: "没有权限!",
			})
		} else {
			core.ShowErrorPage(c, 403, "没有权限!")
		}

		return nil
	}
}

func ReloadUserAccessPrivileges() {

	if keys := userRoleMap.Keys(); len(keys) > 0 {
		for _, key := range keys {
			userRoleMap.Remove(key)
		}
	}
	if keys := userPrivMap.Keys(); len(keys) > 0 {
		for _, key := range keys {
			userPrivMap.Remove(key)
		}
	}

	if users := services.AdminService.GetAdminUsers(); users != nil {
		for _, user := range users {
			userRoleMap.Set(strconv.FormatInt(user.ID, 10), user.Role)
			if user.Role != SuperUserRole {
				if role := services.AdminService.GetAdminRoleById(user.Role); role != nil {
					if len(role.Privileges) > 0 {
						ps := SplitToInt64Slice(role.Privileges, ",")
						for _, p := range ps {
							if p > 0 {
								if privilege := services.AdminService.GetAdminPrivilegeById(p); privilege != nil {
									k := strings.ToLower(fmt.Sprintf("%d %s", user.ID, privilege.Privilege))
									userPrivMap.Set(k, "1")
								}
							}
						}
					}
				}
			}
		}
	}
}
