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

type MemoryApp struct{}

func (m *MemoryApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "memory",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e memory.tar.gz ] && tar -xzvf memory.tar.gz || ( curl -OJ http://" + config.Config().HttpServer.Addr + "/api/v1/download/memory.tar.gz && tar -xzvf memory.tar.gz)",
		Prepare:       "cd /tmp/tune/memory && sh prepare.sh",
		Tune:          "cd /tmp/tune/memory && atune-adm tuning --project stream --detail tuning_stream_client.yaml",
		Restore:       "cd /tmp/tune/memory && atune-adm tuning --restore --project stream",
	}
	return info
}
