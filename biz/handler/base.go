package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
)

// 状态码常量
const (
	SuccessMsg = "success"

	SuccessCode      = 0
	ServiceErrCode   = 10001
	ParamErrCode     = 10002
	CompileErrCode   = 20001
	RuleNotExistCode = 20002
	RuleExecErrCode  = 20003
)

// BaseResp 统一 Http 响应格式
type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BindResp(c *app.RequestContext, code int, msg string, data interface{}) {
	c.JSON(200, BaseResp{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
