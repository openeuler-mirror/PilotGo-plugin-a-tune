/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type KeyParametersSelectApp struct{}

func (kps *KeyParametersSelectApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "key_parameters_select",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e key_parameters_select.tar.gz ] && tar -xzvf key_parameters_select.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/key_parameters_select.tar.gz && tar -xzvf key_parameters_select.tar.gz)",
		Prepare:       "cd /tmp/tune/key_parameters_select && sh prepare.sh",
		Tune:          "cd /tmp/tune/key_parameters_select && atune-adm tuning --project key_parameters_select --detail key_parameters_select_client.yaml",
		Restore:       "cd /tmp/tune/key_parameters_select && atune-adm tuning --restore --project key_parameters_select",
	}
	return info
}
