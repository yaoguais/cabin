package admin

import (
	"bytes"
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/core/middlewares"
	"github.com/Yaoguais/gadmin/core/models"
	"github.com/Yaoguais/gadmin/core/services"
	. "github.com/Yaoguais/gadmin/lib/db"
	"github.com/Yaoguais/gadmin/lib/json"
	"github.com/Yaoguais/gadmin/lib/model"
	. "github.com/Yaoguais/gadmin/lib/slice"
	. "github.com/Yaoguais/gadmin/lib/string"
	tpl "github.com/Yaoguais/gadmin/lib/template"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var adminUserModelDtConfig *datatable.DtConfig
var adminRoleModelDtConfig *datatable.DtConfig
var adminPrivilegeModelDtConfig *datatable.DtConfig

func init() {
	adminUserModelDtConfig = datatable.NewDtConfig(models.AdminUserModel{})
	adminRoleModelDtConfig = datatable.NewDtConfig(models.AdminRoleModel{})
	adminRoleModelDtConfig.RowCB = GetRoleListRowCB
	adminPrivilegeModelDtConfig = datatable.NewDtConfig(models.AdminPrivilegeModel{})
}

type AdminController struct {
}

func (*AdminController) GetUserList(c echo.Context) error {
	roles := services.AdminService.GetAdminRoles()
	d := tpl.Vars{
		"DtConf":    template.HTML(json.JsonEncode(adminUserModelDtConfig)),
		"Roles":     roles,
		"RolesJson": json.JsonEncode(roles),
	}

	return core.Render(c, core.GetCoreViewPath("admin/users.html"), d)
}

func (*AdminController) PostUserDataTable(c echo.Context) error {
	result := datatable.GetModelDtData(c, adminUserModelDtConfig)

	return c.JSON(http.StatusOK, result)
}

func (*AdminController) PostCreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	role, _ := strconv.ParseInt(c.FormValue("role"), 10, 64)

	if len(username) == 0 || len(password) == 0 || role <= 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	if u := services.AdminService.GetAdminUserByUsername(username); u != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "用户名已存在!",
		})
	}

	_, err := Db.Exec("INSERT INTO admin_users(username,password,role) VALUES(?,?,?)",
		username,
		services.AdminService.MakePassword(password),
		role,
	)

	if err != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "系统繁忙, 请重试!",
		})
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "添加成功!",
	})
}

func (*AdminController) PostDeleteUser(c echo.Context) error {
	values, _ := c.FormParams()
	for k := range values {
		if r := len(k) - 1; r > 3 {
			ids := k[3:r]
			if id, err := strconv.ParseInt(ids, 10, 64); err == nil && id > 0 && id != middlewares.SuperUserId {
				services.AdminService.DeleteAdminUserByUid(id)
			}
		}
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "操作成功!",
	})
}

func (*AdminController) GetEditUser(c echo.Context) error {
	ids := c.QueryParam("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil || id <= 0 {
		return core.ShowErrorPage(c, 403, "没有权限!")
	}

	u := services.AdminService.GetAdminUserByUid(id)
	if u == nil {
		return core.ShowErrorPage(c, 401, "参数错误!")
	}

	roles := services.AdminService.GetAdminRoles()
	d := tpl.Vars{
		"Roles": roles,
		"User":  u,
	}

	return core.Render(c, core.GetCoreViewPath("admin/edit_user.html"), d)
}

func (*AdminController) PostEditUser(c echo.Context) error {
	ids := c.FormValue("id")
	username := c.FormValue("username")
	password := c.FormValue("password")
	role, _ := strconv.ParseInt(c.FormValue("role"), 10, 64)

	if len(username) == 0 || role <= 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil || id <= 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	u := services.AdminService.GetAdminUserByUid(id)
	if u == nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	if u.Username != username {
		if t := services.AdminService.GetAdminUserByUsername(username); t != nil {
			return core.Response(c, &core.Resp{
				Ret: 1,
				Msg: "用户名已存在!",
			})
		}
	}

	if u.ID == middlewares.SuperUserId {
		role = middlewares.SuperUserRole
	}

	newUser := &models.AdminUserModel{
		ID:       id,
		Username: username,
		Password: password,
		Role:     role,
	}

	if err := services.AdminService.UpdateUser(newUser); err != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "服务器繁忙!",
		})
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "修改成功!",
	})
}

func (*AdminController) GetRoleList(c echo.Context) error {
	d := tpl.Vars{
		"DtConf":     template.HTML(json.JsonEncode(adminRoleModelDtConfig)),
		"Privileges": services.AdminService.GetAdminPrivileges(),
	}

	return core.Render(c, core.GetCoreViewPath("admin/roles.html"), d)
}

func GetRoleListRowCB(row []interface{}) {
	var pb bytes.Buffer

	if p, ok := row[3].(string); ok {
		ps := strings.Split(p, ",")
		for _, v := range ps {
			if id, err := strconv.ParseInt(v, 10, 64); err == nil && id > 0 {
				if pr := services.AdminService.GetAdminPrivilegeById(id); pr != nil {
					if pb.Len() > 0 {
						pb.WriteString("&nbsp;&nbsp;")
					}
					pb.WriteString(pr.Name)
				}
			}
		}
		row[3] = pb.String()
	}
}

func (*AdminController) PostRoleDataTable(c echo.Context) error {
	result := datatable.GetModelDtData(c, adminRoleModelDtConfig)

	return c.JSON(http.StatusOK, result)
}

func (*AdminController) PostCreateRole(c echo.Context) error {
	name := c.FormValue("name")
	form := c.Request().Form
	var privileges string
	if ps, ok := form["privileges"]; ok {
		privileges = strings.Join(ps, ",")
	}

	if len(name) == 0 || len(privileges) == 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	_, err := Db.Exec("INSERT INTO admin_roles(name,privileges) VALUES(?,?)",
		name,
		privileges,
	)

	if err != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "系统繁忙, 请重试!",
		})
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "添加成功!",
	})
}

func (*AdminController) PostDeleteRole(c echo.Context) error {
	values, _ := c.FormParams()
	for k := range values {
		if r := len(k) - 1; r > 3 {
			ids := k[3:r]
			if id, err := strconv.ParseInt(ids, 10, 64); err == nil && id > 0 && id != middlewares.SuperUserRole {
				services.AdminService.DeleteAdminRoleById(id)
			}
		}
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "操作成功!",
	})
}

func (*AdminController) GetEditRole(c echo.Context) error {
	ids := c.QueryParam("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil || id <= 0 {
		return core.ShowErrorPage(c, 403, "没有权限!")
	}

	r := services.AdminService.GetAdminRoleById(id)
	if r == nil {
		return core.ShowErrorPage(c, 401, "参数错误!")
	}

	var ps []map[string]interface{}
	privileges := services.AdminService.GetAdminPrivileges()
	rolePs := SplitToInt64Slice(r.Privileges, ",")
	for _, p := range privileges {
		v := map[string]interface{}{
			"ID":   p.ID,
			"Name": p.Name,
		}
		if r.ID == middlewares.SuperUserRole || InInt64Array(p.ID, rolePs) {
			v["Selected"] = true
		} else {
			v["Selected"] = false
		}
		ps = append(ps, v)
	}

	d := tpl.Vars{
		"Privileges": ps,
		"Role":       r,
	}

	return core.Render(c, core.GetCoreViewPath("admin/edit_role.html"), d)
}

func (*AdminController) PostEditRole(c echo.Context) error {
	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	name := c.FormValue("name")
	form := c.Request().Form
	var privileges string
	if ps, ok := form["privileges"]; ok {
		privileges = strings.Join(ps, ",")
	}

	if id == middlewares.SuperUserRole {
		return core.Response(c, &core.Resp{
			Ret: 0,
			Msg: "修改成功!",
		})
	}

	if len(name) == 0 || len(privileges) == 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	_, err := Db.Exec("UPDATE admin_roles SET name = ?, privileges = ? WHERE id = ?",
		name,
		privileges,
		id,
	)

	if err != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "系统繁忙, 请重试!",
		})
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "修改成功!",
	})
}

func (*AdminController) GetPrivilegeList(c echo.Context) error {
	d := tpl.Vars{
		"DtConf": template.HTML(json.JsonEncode(adminPrivilegeModelDtConfig)),
	}

	return core.Render(c, core.GetCoreViewPath("admin/privileges.html"), d)
}

func (*AdminController) PostPrivilegeDataTable(c echo.Context) error {
	result := datatable.GetModelDtData(c, adminPrivilegeModelDtConfig)

	return c.JSON(http.StatusOK, result)
}

func (*AdminController) PostCreatePrivilege(c echo.Context) error {
	name := c.FormValue("name")
	group := c.FormValue("group")
	privilege := c.FormValue("privilege")

	if len(name) == 0 || len(group) == 0 || len(privilege) == 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	_, err := Db.Exec("INSERT INTO admin_privileges(name,`group`,privilege) VALUES(?,?,?)",
		name,
		group,
		privilege,
	)

	if err != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "系统繁忙, 请重试!",
		})
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "添加成功!",
	})
}

func (*AdminController) PostDeletePrivilege(c echo.Context) error {
	values, _ := c.FormParams()
	for k := range values {
		if r := len(k) - 1; r > 3 {
			ids := k[3:r]
			if id, err := strconv.ParseInt(ids, 10, 64); err == nil && id > 0 && id != middlewares.SuperUserRole {
				services.AdminService.DeleteAdminPrivilegeById(id)
			}
		}
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "操作成功!",
	})
}

func (*AdminController) GetEditPrivilege(c echo.Context) error {
	id, _ := strconv.ParseInt(c.QueryParam("id"), 10, 64)

	p := services.AdminService.GetAdminPrivilegeById(id)
	if p == nil {
		return core.ShowErrorPage(c, 401, "参数错误!")
	}

	d := tpl.Vars{
		"Privilege": p,
	}

	return core.Render(c, core.GetCoreViewPath("admin/edit_privilege.html"), d)
}

func (*AdminController) PostEditPrivilege(c echo.Context) error {
	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	name := c.FormValue("name")
	group := c.FormValue("group")
	privilege := c.FormValue("privilege")

	if len(name) == 0 || len(group) == 0 || len(privilege) == 0 {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "参数不合法!",
		})
	}

	_, err := Db.Exec("UPDATE admin_privileges SET name = ?, `group` = ?, privilege = ? WHERE id = ?",
		name,
		group,
		privilege,
		id,
	)

	if err != nil {
		return core.Response(c, &core.Resp{
			Ret: 1,
			Msg: "系统繁忙, 请重试!",
		})
	}

	middlewares.ReloadUserAccessPrivileges()

	return core.Response(c, &core.Resp{
		Ret: 0,
		Msg: "修改成功!",
	})
}
