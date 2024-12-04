/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type KafkaApp struct{}

type KafkaImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (k *KafkaApp) Info() *KafkaImp {
	info := &KafkaImp{
		BaseTune: TuneInfo{
			TuneName:      "kafka",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e kafka.tar.gz ] && tar -xzvf kafka.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/kafka.tar.gz && tar -xzvf kafka.tar.gz)",
			Prepare:       "cd /tmp/tune/kafka && sh prepare.sh",
			Tune:          "cd /tmp/tune/kafka && atune-adm tuning --project kafka --detail ./kafka_client.yaml",
			Restore:       "cd /tmp/tune/kafka && atune-adm tuning --restore --project kafka",
		},
		Notes: "使用kafka atune之前, 请先配置好kafka环境",
	}
	return info
}
