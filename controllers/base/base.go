package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	//result Result // 返回值
	//valid   = validation.Validation{}
	// err     error
	// isValid bool
)

type HttpCode int

type JsonResponse struct {
	Code    HttpCode    `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Prepare() {
	logs.Info("[Prepare]")
}

func (c *ApiController) JsonReturn(message string, result interface{}, code int) {
	c.Data["json"] = JsonResponse{
		Message: message,
		Result:  result,
		Code:    HttpCode(code),
	}
	c.ServeJSON(true)
}
