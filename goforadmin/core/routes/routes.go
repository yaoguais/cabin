package routes

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/core/controllers/admin"
	"github.com/Yaoguais/gadmin/core/controllers/component"
	"github.com/Yaoguais/gadmin/core/controllers/index"
	"github.com/Yaoguais/gadmin/core/middlewares"
	"github.com/labstack/echo/middleware"
)

func RegisterRoutes() {
	app := core.GetApp()
	indexController := index.IndexController{}
	loginController := index.LoginController{}
	logoutController := index.LogoutController{}
	adminController := admin.AdminController{}
	componentController := component.ComponentController{}

	app.Use(middleware.Recover())
	middlewares.ReloadUserAccessPrivileges()

	app.GET("/login", loginController.GetLogin)
	app.POST("/login", loginController.PostLogin)

	a := app.Group("", middlewares.Auth, middlewares.Admin)
	a.GET("/admin/users", adminController.GetUserList)
	a.POST("/admin/users/datatable", adminController.PostUserDataTable)
	a.POST("/admin/users/create", adminController.PostCreateUser)
	a.POST("/admin/users/delete", adminController.PostDeleteUser)
	a.GET("/admin/users/edit", adminController.GetEditUser)
	a.POST("/admin/users/edit", adminController.PostEditUser)
	a.GET("/admin/roles", adminController.GetRoleList)
	a.POST("/admin/roles/datatable", adminController.PostRoleDataTable)
	a.POST("/admin/roles/create", adminController.PostCreateRole)
	a.POST("/admin/roles/delete", adminController.PostDeleteRole)
	a.GET("/admin/roles/edit", adminController.GetEditRole)
	a.POST("/admin/roles/edit", adminController.PostEditRole)
	a.GET("/admin/privileges", adminController.GetPrivilegeList)
	a.POST("/admin/privileges/datatable", adminController.PostPrivilegeDataTable)
	a.POST("/admin/privileges/create", adminController.PostCreatePrivilege)
	a.POST("/admin/privileges/delete", adminController.PostDeletePrivilege)
	a.GET("/admin/privileges/edit", adminController.GetEditPrivilege)
	a.POST("/admin/privileges/edit", adminController.PostEditPrivilege)
	a.GET("/component", componentController.GetIndex)
	a.POST("/component/selects/ajax", componentController.PostSelectsAjax)

	g := app.Group("", middlewares.Auth, middlewares.Access)
	g.GET("/", indexController.GetIndex)
	g.GET("/logout", logoutController.GetLogout)
	g.GET("/resource/token", componentController.GetQiniuResourceToken)
}
