/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

import (
	"openeuler.org/PilotGo/atune-plugin/config"
)

type GraphicsmagickApp struct{}

func (g *GraphicsmagickApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "graphicsmagick",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e graphicsmagick.tar.gz ] && tar -xzvf graphicsmagick.tar.gz || ( curl -OJ http://" + config.Config().HttpServer.Addr + "/api/v1/download/graphicsmagick.tar.gz && tar -xzvf graphicsmagick.tar.gz)",
		Prepare:       "cd /tmp/tune/graphicsmagick && sh prepare.sh",
		Tune:          "cd /tmp/tune/graphicsmagick && atune-adm tuning --project graphicsmagick  --detail gm_client.yaml",
		Restore:       "cd /tmp/tune/graphicsmagick && atune-adm tuning --restore --project graphicsmagick",
	}
	return info
}
