/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type RedisApp struct{}

type RedisImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (redis *RedisApp) Info() *RedisImp {
	info := &RedisImp{
		BaseTune: TuneInfo{
			TuneName:      "redis",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e redis.tar.gz ] && tar -xzvf redis.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/redis.tar.gz && tar -xzvf redis.tar.gz)",
			Prepare:       "cd /tmp/tune/redis && sh prepare.sh",
			Tune:          "cd /tmp/tune/redis && atune-adm tuning --project redis --detail ./redis_client.yaml",
			Restore:       "cd /tmp/tune/redis && atune-adm tuning --restore --project redis",
		},
		Notes: "若要启动基准测试, 在调优前先进入调优工作目录/tmp/tune/redis执行“sh redis_benchmark.sh”,本地主机将访问基准测试主机并触发它。基准测试主机将在基准测试之后将日志文件传输到localhost。",
	}
	return info
}
