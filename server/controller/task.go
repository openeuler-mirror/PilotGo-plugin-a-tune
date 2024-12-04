/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package controller

import (
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func StartTask(c *gin.Context) {
	d := &struct {
		TuneID int `json:"tuneId"`
		TaskID int `json:"taskId"`
	}{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}

	machine_uuids, err := dao.GetResultUUIDByTaskId(strconv.Itoa(d.TaskID))
	if err != nil || len(machine_uuids) == 0 {
		response.Fail(c, nil, "未找到相应的机器")
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
	err = service.RunTask(d.TaskID, d.TuneID, commands, machine_uuids)
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
