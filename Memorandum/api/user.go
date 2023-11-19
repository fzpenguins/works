package api

import (
	"Memorandum/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// Register @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param body body model.User true "新用户信息"
// @Success 200 {object} serialize.Response "{"status":200,"data":{},"msg":"注册成功"}"
// @Failure 400  {object} serialize.Response "{"status":400,"data":{},"Msg":"输入不合规定"}"
// @Failure 500  {object} serialize.Response "{"status":500,"data":{},"Msg":"注册成功"}"
// @Router /user/register [post]
func Register(c context.Context, ctx *app.RequestContext) {
	res := service.Register(ctx)
	if res.Status != 200 {
		ctx.JSON(int(res.Status), res)
	} else {
		ctx.JSON(int(res.Status), res)
	}
}

// Login @Tags USER
// @Summary 用户登录
// @Produce json
// @Accept json
// @Param body body model.User true "用户信息"
// @Success 200 {object} serialize.Response "{"status":200,"data":serialize.TokenData,"msg":"登录成功"}"
// @Failure 400 {object} serialize.Response "{"status":400,"data":{},"Msg":"输入不合规定"}|"
// @Failure 400 {object} serialize.Response "{"status":400,"data":{},"Msg":"验证失败,发生错误"}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"Msg":"查无此号"}"
// @Router /user/login [post]
func Login(c context.Context, ctx *app.RequestContext) {
	res := service.Login(ctx)
	if res.Status != 200 {
		ctx.JSON(int(res.Status), res)
	} else {
		ctx.JSON(int(res.Status), res)
	}
}
