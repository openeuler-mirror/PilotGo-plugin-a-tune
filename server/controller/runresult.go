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
