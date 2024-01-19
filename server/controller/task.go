package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func StartTask(c *gin.Context) {
	d := &struct {
		MachineUUIDs []string `json:"machine_uuids"`
		TuneID       int      `json:"tune_id"`
		TaskID       int      `json:"taskId"`
	}{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}
	commands, err := service.GetCommandByID(d.TuneID)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	err = service.UpdateTaskStatus(d.TaskID, service.TaskRunnig)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	err = service.RunTask(d.TaskID, d.TuneID, commands, d.MachineUUIDs)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已启动执行任务")
}

func CreatTask(c *gin.Context) {
	d := &struct {
		MachineUUIDs []string `json:"machine_uuids"`
		TuneID       int      `json:"tune_id"`
		TaskName     string   `json:"task_name"`
		Description  string   `json:"description"`
	}{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}
	err := service.SaveTask(d.TaskName, d.Description, d.MachineUUIDs, d.TuneID)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "指令下发完成")
}

func TaskLists(c *gin.Context) {
	query := &response.PaginationQ{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.QueryTaskLists(query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	response.DataPagination(c, data, total, query)
}

func DeleteTask(c *gin.Context) {
	taskdel := struct {
		TaskID []int `json:"ids"`
	}{}
	if err := c.Bind(&taskdel); err != nil {
		response.Fail(c, nil, "parameter error")
		return
	}

	if err := service.DeleteTask(taskdel.TaskID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已删除任务")
}

func SearchTask(c *gin.Context) {
	search := c.Query("search")

	query := &response.PaginationQ{}
	if err := c.ShouldBindQuery(query); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}

	data, total, err := service.SearchTask(search, query)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.DataPagination(c, data, total, query)
}
