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

type FioApp struct{}

type FioImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (f *FioApp) Info() *FioImp {
	info := &FioImp{
		BaseTune: TuneInfo{
			TuneName:      "fio",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e fio.tar.gz ] && tar -xzvf fio.tar.gz || ( curl -OJ http://" + config.Config().HttpServer.Addr + "/api/v1/download/fio.tar.gz && tar -xzvf fio.tar.gz)",
			Prepare:       "cd /tmp/tune/fio && sh prepare.sh 参数一 参数二",
			Tune:          "cd /tmp/tune/fio && atune-adm tuning --project fio --detail tuning_fio_client.yaml",
			Restore:       "cd /tmp/tune/fio && atune-adm tuning --restore --project fio",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [test_path] [test_disk]",
	}
	return info
}
