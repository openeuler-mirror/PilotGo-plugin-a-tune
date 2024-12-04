/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type SparkApp struct{}

type SparkImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (spark *SparkApp) Info() *SparkImp {
	info := &SparkImp{
		BaseTune: TuneInfo{
			TuneName:      "spark",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e spark.tar.gz ] && tar -xzvf spark.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/spark.tar.gz && tar -xzvf spark.tar.gz)",
			Prepare:       "cd /tmp/tune/spark && sh run_env.sh",
			Tune:          "",
			Restore:       "",
		},
		Notes: "调优步骤均在run_env.sh完成, 执行该脚本可完成环境的部署和调优测试",
	}
	return info
}
