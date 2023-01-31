package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/trust/easy_rule_engine/biz/dal"
	"strconv"
)

// HandleAddExpression 增加表达式, 存储在 DB , 返回 Id , 如果数据库已经存在该表达式, 则直接返回 Id
func HandleAddExpression(ctx context.Context, c *app.RequestContext) {
	type ExpReq struct {
		Exp string `json:"exp"`
	}
	req := ExpReq{}
	err := c.Bind(&req)
	if err != nil || len(req.Exp) == 0 {
		BindResp(c, ParamErrCode, err.Error(), nil)
		return
	}
	d, err := dal.AddExpression(req.Exp)
	if err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	BindResp(c, SuccessCode, SuccessMsg, d)
}

// HandleDeleteExpression 删除表达式
func HandleDeleteExpression(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 64)
	if err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	// 查询被删除的表达式
	exp, err := dal.GetExpressionByID(uint(id))
	if err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	// 不存在该 id 则返回错误信息
	if exp.ID == 0 {
		BindResp(c, RuleNotExistCode, "exp id "+c.Param("id")+" not exist", nil)
		return
	}
	// 执行删除
	err = dal.DeleteExpressionByID(uint(id))
	if err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	// 删除成功返回被删除的表达式
	BindResp(c, SuccessCode, SuccessMsg, exp)
}

// HandleGetAllExpression 查询所有表达式
func HandleGetAllExpression(ctx context.Context, c *app.RequestContext) {
	exps, err := dal.GetAllExpression()
	if err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	BindResp(c, SuccessCode, SuccessMsg, exps)
}

// HandleRunExpression 运行指定表达式, 请求参数传入表达式 Id 以及表达式中指定的参数
func HandleRunExpression(ctx context.Context, c *app.RequestContext) {
	type ExpReq struct {
		ExpId  uint
		Params map[string]interface{}
	}
	var expReq ExpReq
	if err := c.Bind(&expReq); err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	// 查询表达式
	exp, err := dal.GetExpressionByID(expReq.ExpId)
	if err != nil {
		BindResp(c, ServiceErrCode, err.Error(), nil)
		return
	}
	// 该 ID 的表达式不存在
	if exp.ID == 0 {
		BindResp(c, RuleNotExistCode, "exp id "+strconv.Itoa(int(expReq.ExpId))+" does not exists", nil)
		return
	}
	// 编译表达式, 得到语法树的根节点
	root, err := Compiler(exp.Exp)
	if err != nil {
		BindResp(c, CompileErrCode, err.Error(), nil)
		return
	}
	// 参数注入
	err = root.Eval(expReq.Params)
	if err != nil {
		BindResp(c, RuleExecErrCode, err.Error(), nil)
		return
	}
	resp, _ := root.GetVal()
	BindResp(c, SuccessCode, SuccessMsg, resp)
}
