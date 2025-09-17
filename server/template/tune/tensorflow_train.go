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

type TensorflowTrainApp struct{}

func (tensor *TensorflowTrainApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "tensorflow_train",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e tensorflow_train.tar.gz ] && tar -xzvf tensorflow_train.tar.gz || ( curl -OJ http://" + config.Config().HttpServer.Addr + "/api/v1/download/tensorflow_train.tar.gz && tar -xzvf tensorflow_train.tar.gz)",
		Prepare:       "cd /tmp/tune/tensorflow_train && sh prepare.sh",
		Tune:          "cd /tmp/tune/tensorflow_train && atune-adm tuning --project tensorflow_train --detail tensorflow_train_client.yaml",
		Restore:       "cd /tmp/tune/tensorflow_train && atune-adm tuning --restore --project tensorflow_train",
	}
	return info
}
