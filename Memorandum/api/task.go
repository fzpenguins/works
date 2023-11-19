package api

import (
	"Memorandum/serialize"
	"Memorandum/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

// CreateTask  @Tags TASK
// @Summary 创建任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Param body body service.CreateTask true "创建任务"
// @Success 200 {object} serialize.Response "{"status":200,"data":model.Task,"msg":"创建成功!"}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"Msg":"创建文件失败!"}"
// @Failure 400 {string} string "参数输入错误"
// @Router /task [post]
func CreateTask(c context.Context, ctx *app.RequestContext) {
	var ct service.CreateTask
	if err := ctx.BindJSON(&ct); err != nil {
		ctx.JSON(400, "参数输入错误")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))

	res := ct.CreateTask(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

// UpdateTask @Tags TASK
// @Summary 更新任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Param body body service.UpdateTask true "任务更新参数"
// @Success 200 {object} serialize.Response "{"status":200,"data":{},"msg":"更新成功"}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"Msg":"查无此事项"}"
// @Failure 400 {string} string "参数输入错误"
// @Router /task [put]
func UpdateTask(c context.Context, ctx *app.RequestContext) {
	var ut service.UpdateTask
	if err := ctx.BindJSON(&ut); err != nil {
		ctx.JSON(400, "参数输入错误")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := ut.UpdateTask(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

// DeleteTask @Tags TASK
// @Summary 删除任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Param id path integer true "任务序号"
// @Success 200 {object} serialize.Response "{"status":200,"data":{},"msg":"事项删除成功"}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"Msg":"所删除事项不存在"}"
// @Failure 400 {string} string "id参数输入有误"
// @Router /task/:id [delete]
func DeleteTask(c context.Context, ctx *app.RequestContext) {
	var err error
	var dt service.DeleteTask
	dt.Id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, "id参数输入有误!")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := dt.DeleteTask(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

// DeleteTasksInStatus @Tags TASK
// @Summary 删除任务指定状态事项
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Param status path integer true "任务状态，0代表查看待办事项，1代表查看已完成事项"
// @Success 200 {object} serialize.Response "{"status":200,"data":{},"msg":"删除成功"}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"Msg":"无可删除事项"}"
// @Failure 400 {string} string "status输入有误"
// @Router /tasks/:status [delete]
func DeleteTasksInStatus(c context.Context, ctx *app.RequestContext) {
	var err error
	var dt service.DeleteTasksInStatus
	dt.Status, err = strconv.Atoi(ctx.Param("status"))
	if err != nil {
		ctx.JSON(400, "status输入有误!")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := dt.DeleteTasksInStatus(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

// DeleteAllTasks  @Tags TASK
// @Summary 删除所有任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} serialize.Response "{"status":200,"data":{},"msg":"删除成功"}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"Msg":"无可删除事项"}"
// @Router /tasks [delete]
func DeleteAllTasks(c context.Context, ctx *app.RequestContext) {
	var dt service.DeleteAllTasks
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := dt.DeleteAllTasks(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

// ShowTasksInStatus   @Tags TASK
// @Summary 查询指定状态任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Param status path integer true "任务状态，0代表查看待办事项，1代表查看已完成事项,其余代表查看所有事项"
// @Success 200 {object} serialize.Response "{"status":200,"data":Data:[]model.Task,"msg":"删除成功"}"
// @Failure 400 {string} string "status参数输入错误"
// @Success 500 {object} serialize.Response "{"status":500,"data":{},"msg":"无"}"
// @Router /tasks/:status [get]
func ShowTasksInStatus(c context.Context, ctx *app.RequestContext) {
	var st service.ShowTasksInStatus
	var sta service.ShowAllTasks
	var err error
	var res serialize.Response
	st.Status, err = strconv.Atoi(ctx.Param("status"))
	if err != nil {
		ctx.JSON(400, "status参数输入错误")
		//ctx.JSON(400, st.Status)
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	if st.Status != 0 && st.Status != 1 {
		res = sta.ShowAllTasks(claim.Id)
	} else {
		res = st.ShowTasksInStatus(claim.Id)
	}
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func ShowAllTasks(c context.Context, ctx *app.RequestContext) {
	var st service.ShowAllTasks
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := st.ShowAllTasks(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

// ShowTasksByKey  @Tags TASK
// @Summary 关键词查询
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param Authorization header string true "访问令牌"
// @Param body body service.ShowTasksByKey true "搜索关键字"
// @Success 200 {object} serialize.Response "{"status":200,"data":[]model.Task,"msg":"删除成功"}"
// @Failure 400 {object} map[string]interface{} "{"err":err}"
// @Failure 500 {object} serialize.Response "{"status":500,"data":{},"msg":"无"}"
// @Router /tasks [get]
func ShowTasksByKey(c context.Context, ctx *app.RequestContext) {
	var st service.ShowTasksByKey
	//st.Key = ctx.Query("key")
	if err := ctx.BindJSON(&st); err != nil {
		ctx.JSON(400, utils.H{
			"err": err,
		})
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := st.ShowTasksByKey(int(claim.Id))
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}
