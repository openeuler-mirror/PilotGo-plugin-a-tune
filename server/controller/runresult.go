package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func QueryResults(c *gin.Context) {
	taskId := c.Query("taskId")

	data, err := service.SearchResultByTaskId(taskId)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, data, "获取到结果详情")
}

func DeleteResult(c *gin.Context) {
	resultdel := struct {
		ResultID []int `json:"ids"`
	}{}
	if err := c.Bind(&resultdel); err != nil {
		response.Fail(c, nil, "parameter error")
		return
	}

	if err := service.DeleteResult(resultdel.ResultID); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "已删除")
}

func QueryTuneResults(c *gin.Context) {
	taskId := c.Query("taskId")
	machine_uuid := c.Query("uuid")

	data, err := service.GetResultAnalisis(taskId, machine_uuid)
	if err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, data, "获取到调优分析详情")
}
