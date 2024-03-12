package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type TensorflowTrainApp struct{}

func (tensor *TensorflowTrainApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "tensorflow_train",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e tensorflow_train.tar.gz ] && tar -xzvf tensorflow_train.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/tensorflow_train.tar.gz && tar -xzvf tensorflow_train.tar.gz)",
		Prepare:       "cd /tmp/tune/tensorflow_train && sh prepare.sh",
		Tune:          "cd /tmp/tune/tensorflow_train && atune-adm tuning --project tensorflow_train --detail tensorflow_train_client.yaml",
		Restore:       "cd /tmp/tune/tensorflow_train && atune-adm tuning --restore --project tensorflow_train",
	}
	return info
}
