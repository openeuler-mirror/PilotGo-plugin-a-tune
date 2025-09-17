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

type MysqlSysbenchApp struct{}

type MysqlSysbenchImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (m *MysqlSysbenchApp) Info() *MysqlSysbenchImp {
	info := &MysqlSysbenchImp{
		BaseTune: TuneInfo{
			TuneName:      "mysql_sysbench",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e mysql_sysbench.tar.gz ] && tar -xzvf mysql_sysbench.tar.gz || ( curl -OJ http://" + config.Config().HttpServer.Addr + "/api/v1/download/mysql_sysbench.tar.gz && tar -xzvf mysql_sysbench.tar.gz)",
			Prepare:       "cd /tmp/tune/mysql_sysbench && sh prepare.sh",
			Tune:          "cd /tmp/tune/mysql_sysbench && atune-adm tuning --project mysql_sysbench --detail mysql_sysbench_client.yaml",
			Restore:       "cd /tmp/tune/mysql_sysbench && atune-adm tuning --restore --project mysql_sysbench",
		},
		Notes: "1,根据指南安装mysql: https://blog.csdn.net/weixin_43214408/article/details/116895091 \n 2,根据指南安装sysbench: https://blog.csdn.net/weixin_43214408/article/details/116898751",
	}
	return info
}
