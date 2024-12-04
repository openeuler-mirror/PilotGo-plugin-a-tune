/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package controller

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/plugin"
	"openeuler.org/PilotGo/atune-plugin/service"
)

func AtuneClientInstall(c *gin.Context) {
	d := &struct {
		MachineUUIDs []string `json:"uuids"`
	}{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}

	run_result := func(result []*common.CmdResult) {
		for _, res := range result {
			logger.Info("结果：%v", *res)
			if err := service.AtuneManage(res, service.CommandInstall_Type); err != nil {
				logger.Error("处理结果失败：%v", err.Error())
			}
		}
	}

	dd := &common.Batch{
		MachineUUIDs: d.MachineUUIDs,
	}
	err := plugin.GlobalClient.RunCommandAsync(dd, service.CommandInstall_Cmd, run_result)
	if err != nil {
		logger.Error("远程调用失败：%v", err.Error())
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "指令下发完成")
}
func AtuneClientRemove(c *gin.Context) {
	d := &struct {
		MachineUUIDs []string `json:"uuids"`
	}{}
	if err := c.ShouldBind(d); err != nil {
		logger.Debug("绑定批次参数失败：%s", err)
		response.Fail(c, nil, "parameter error")
		return
	}

	run_result := func(result []*common.CmdResult) {
		for _, res := range result {
			logger.Info("结果：%v", *res)
			if err := service.AtuneManage(res, service.CommandUninstall_Type); err != nil {
				logger.Error("处理结果失败：%v", err.Error())
			}
		}
	}

	dd := &common.Batch{
		MachineUUIDs: d.MachineUUIDs,
	}
	err := plugin.GlobalClient.RunCommandAsync(dd, service.CommandUninstall_Cmd, run_result)
	if err != nil {
		logger.Error("远程调用失败：%v", err.Error())
		response.Fail(c, nil, err.Error())
		return
	}
	response.Success(c, nil, "指令下发完成")
}
