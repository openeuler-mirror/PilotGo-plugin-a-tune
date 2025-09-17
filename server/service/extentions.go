/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
)

// GetExtentions returns the plugin extentions for etcd registration
func GetExtentions() []common.Extention {
	var ex []common.Extention
	me1 := &common.MachineExtention{
		Type:       common.ExtentionMachine,
		Name:       "安装a-tune",
		URL:        "/plugin/atune/atune_install",
		Permission: "plugin.atune.agent/install",
	}
	me2 := &common.MachineExtention{
		Type:       common.ExtentionMachine,
		Name:       "卸载a-tune",
		URL:        "/plugin/atune/atune_uninstall",
		Permission: "plugin.atune.agent/uninstall",
	}
	pe1 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "任务列表",
		URL:        "/task",
		Permission: "plugin.atune.page/menu",
	}
	pe2 := &common.PageExtention{
		Type:       common.ExtentionPage,
		Name:       "调优模板",
		URL:        "/template",
		Permission: "plugin.atune.page/menu",
	}
	ex = append(ex, me1, me2, pe1, pe2)
	return ex
}
