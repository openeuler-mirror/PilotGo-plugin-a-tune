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
