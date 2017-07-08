package core

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/Yaoguais/gadmin/lib/db"
	"github.com/Yaoguais/gadmin/lib/log"
	"github.com/Yaoguais/gadmin/lib/redis"
	"github.com/Yaoguais/gadmin/lib/session"
	"github.com/Yaoguais/gadmin/lib/template"
	"github.com/labstack/echo"
)

var app *echo.Echo

func init() {
	appRoot = getAppRoot()
	app = echo.New()
}

func GetApp() *echo.Echo {
	return app
}

func RunApp() {
	if err := app.Start(appHost); err != nil {
		log.Fatal("start server failed ", err)
	}
}

func SetConfig(file string) {
	configFile = file
}

func InitApp() {
	initConfig()
	initLog()
	initException()
	initStaticFile()
	initTemplateGlobalVars()
	initDatabase()
	initRedis()
	initSession()
}

func initException() {
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		var code int
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		} else {
			code = http.StatusInternalServerError
		}

		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD {
				if err := c.NoContent(code); err != nil {
					goto ERROR
				}
			} else {
				ShowErrorPage(c, code, "服务器繁忙!")
			}
		}

	ERROR:
		app.Logger.Error(err)
	}
}

func initLog() {
	file := fmt.Sprintf("%s/%s/app-web.log", GetAppRoot(), logPath)
	log.SetOutputPath(file)
	log.SetOutputLevel(logLevel)
}

func initSession() {
	if sessionDriver == "redis" {
		session.SetRedisClient(redis.Redis)
	} else {
		// file
		session.SetSessionPath(GetAppRoot() + "/" + sessionPath)
	}
	session.SetSessionDriver(sessionDriver)
	session.SetSessionSecret(sessionSecret)
	session.StartSession()
}

func initStaticFile() {
	app.Static("/static", GetAppRoot()+"/public/static")
}

func initTemplateGlobalVars() {
	vars := template.Vars{
		"WebsiteTitle": appName,
		"StaticHost":   appStaticHost,
		"layout":       GetCoreViewPath("layout/layout.html"),
	}
	template.RegisterGlobalVars(vars)
}

func initDatabase() {
	db.Connect(dbConf)
}

func initRedis() {
	redis.Connect(redisConf)
}

func getAppRoot() string {
	execFileRelativePath, _ := exec.LookPath(os.Args[0])
	execDirRelativePath, _ := path.Split(execFileRelativePath)
	execDirAbsPath, _ := filepath.Abs(execDirRelativePath)

	return execDirAbsPath
}

func GetAppRoot() string {
	return appRoot
}

func GetAppViewPath(file string) string {
	return appRoot + "/app/views/" + file
}

func GetCoreViewPath(file string) string {
	return appRoot + "/" + coreView + "/" + file
}

func GetAppName() string {
	return appName
}

func GetUserId(c echo.Context) string {
	session := session.GetSession(c, session.SessionName())
	if id, ok := session.Values["id"]; ok {
		if ret, ok := id.(string); ok {
			return ret
		}
	}

	return ""
}

func GetUploadBucket() string {
	return uploadBucket
}

func GetUploadDomain() string {
	return uploadDomain
}
