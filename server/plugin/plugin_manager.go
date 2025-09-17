/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package plugin

import (
	"gitee.com/openeuler/PilotGo/sdk/plugin/client"
	"openeuler.org/PilotGo/atune-plugin/config"
)

var (
	GlobalClient *client.Client
)

func Init(plugin *config.PluginAtune) interface{} {
	// PluginInfo 结构在新版 SDK 中已移除，返回空接口
	return nil
}
