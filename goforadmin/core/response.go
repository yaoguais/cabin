package core

import (
	"fmt"
	"github.com/Yaoguais/gadmin/lib/log"
	"github.com/Yaoguais/gadmin/lib/template"
	"github.com/labstack/echo"
	"net/http"
)

type Resp struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(c echo.Context, data *Resp) error {

	log.Info(fmt.Sprintf("request %s %s response %v",
		c.Request().Method,
		c.Request().RequestURI,
		data),
	)

	return c.JSON(http.StatusOK, data)
}

func ShowErrorPage(c echo.Context, code int, message string) error {
	d := template.Vars{
		"Title":   fmt.Sprintf("%d 错误", code),
		"Message": message,
	}

	log.Info(fmt.Sprintf("request %s %s response message %s http status %d",
		c.Request().Method,
		c.Request().RequestURI,
		message,
		code),
	)

	return RenderWithCode(c, GetCoreViewPath("error/error.html"), d, code)
}

func Render(c echo.Context, file string, data interface{}) error {
	return RenderWithCode(c, file, data, http.StatusOK)
}

func RenderWithCode(c echo.Context, file string, data interface{}, code int) error {
	str, err := template.ParseFileToString(file, data)
	if err != nil {
		str = "模板缺失!"
	}

	log.Info(fmt.Sprintf("request %s %s response http status %d",
		c.Request().Method,
		c.Request().RequestURI,
		code),
	)

	return c.HTML(code, str)
}
