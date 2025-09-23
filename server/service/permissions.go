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

// GetPermissions returns the plugin permissions for etcd registration
func GetPermissions() []common.Permission {
	var pe []common.Permission
	p1 := common.Permission{
		Resource: "atune_operate",
		Operate:  "button",
	}
	p2 := common.Permission{
		Resource: "atune",
		Operate:  "menu",
	}
	p3 := common.Permission{
		Resource: "plugin.atune.agent",
		Operate:  "install",
	}
	p4 := common.Permission{
		Resource: "plugin.atune.agent",
		Operate:  "uninstall",
	}
	return append(pe, p1, p2, p3, p4)
}
