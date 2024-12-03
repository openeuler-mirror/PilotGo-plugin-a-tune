/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type OpenGaussApp struct{}

type OpenGaussImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (gauss *OpenGaussApp) Info() *OpenGaussImp {
	info := &OpenGaussImp{
		BaseTune: TuneInfo{
			TuneName:      "openGauss",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e openGauss.tar.gz ] && tar -xzvf openGauss.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/openGauss.tar.gz && tar -xzvf openGauss.tar.gz)",
			Prepare:       "cd /tmp/tune/openGauss && sh prepare.sh",
			Tune:          "cd /tmp/tune/openGauss && atune-adm tuning --project openGauss_tpcc --detail openGauss_client.yaml",
			Restore:       "cd /tmp/tune/openGauss && atune-adm tuning --restore --project openGauss_tpcc",
		},
		Notes: "请先安装openGauss和benchmarksql-5.0 \n openGauss安装指南: https://docs.opengauss.org/zh/docs/2.1.0/docs/installation/installation.html \n benchmarksql-5.0下载地址: https://udomain.dl.sourceforge.net/project/benchmarksql/benchmarksql-5.0.zip",
	}
	return info
}
