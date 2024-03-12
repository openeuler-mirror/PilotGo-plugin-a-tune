package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type KeyParametersSelectVariantApp struct{}

func (kpsv *KeyParametersSelectVariantApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "key_parameters_select_variant",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e key_parameters_select_variant.tar.gz ] && tar -xzvf key_parameters_select_variant.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/key_parameters_select_variant.tar.gz && tar -xzvf key_parameters_select_variant.tar.gz)",
		Prepare:       "cd /tmp/tune/key_parameters_select_variant && sh prepare.sh",
		Tune:          "cd /tmp/tune/key_parameters_select_variant && atune-adm tuning --project key_parameters_select_variant --detail key_parameters_select_variant_client.yaml",
		Restore:       "cd /tmp/tune/key_parameters_select_variant && atune-adm tuning --restore --project key_parameters_select_variant",
	}
	return info
}
