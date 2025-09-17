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

type CompressExceptApp struct{}

func (ce *CompressExceptApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "compress_Except_example",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e compress_Except_example.tar.gz ] && tar -xzvf compress_Except_example.tar.gz || ( curl -OJ http://" + config.Config().HttpServer.Addr + "/api/v1/download/compress_Except_example.tar.gz  && tar -xzvf compress_Except_example.tar.gz)",
		Prepare:       "cd /tmp/tune/compress_Except_example && sh prepare.sh enwik8.zip",
		Tune:          "cd /tmp/tune/compress_Except_example && atune-adm tuning --project compress_Except_example --detail compress_Except_example_client.yaml",
		Restore:       "cd /tmp/tune/compress_Except_example && atune-adm tuning --restore --project compress_Except_example",
	}
	return info

}
