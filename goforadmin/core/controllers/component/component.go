package component

import (
	"github.com/Yaoguais/gadmin/core"
	"github.com/Yaoguais/gadmin/lib/random"
	tpl "github.com/Yaoguais/gadmin/lib/template"
	"github.com/labstack/echo"
	"github.com/qiniu/api.v7/kodo"
	"net/http"
)

type ComponentController struct {
}

func (*ComponentController) GetIndex(c echo.Context) error {
	d := tpl.Vars{
		"UploadDomain": core.GetUploadDomain(),
	}

	return core.Render(c, core.GetCoreViewPath("component/index.html"), d)
}

func (*ComponentController) PostSelectsAjax(c echo.Context) error {
	keyword := c.FormValue("keyword")

	limit := 10
	var data [][]interface{} = make([][]interface{}, 0, 10)
	for i := 0; i < limit; i++ {
		row := []interface{}{
			i,
			random.RandowmString(4) + keyword + random.RandowmString(2),
		}
		data = append(data, row)
	}

	return core.Response(c, &core.Resp{
		Ret:  0,
		Data: data,
	})
}

func (*ComponentController) GetQiniuResourceToken(c echo.Context) error {
	k := kodo.New(0, nil)
	policy := &kodo.PutPolicy{
		Scope:   core.GetUploadBucket(),
		Expires: 3600,
	}
	token := k.MakeUptoken(policy)

	return c.JSON(http.StatusOK, &map[string]string{
		"uptoken": token,
	})
}
