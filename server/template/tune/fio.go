package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type FioApp struct{}

type FioImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (f *FioApp) Info() *FioImp {
	info := &FioImp{
		BaseTune: TuneInfo{
			TuneName:      "fio",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e fio.tar.gz ] && tar -xzvf fio.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/fio.tar.gz && tar -xzvf fio.tar.gz)",
			Prepare:       "cd /tmp/tune/fio && sh prepare.sh 参数一 参数二",
			Tune:          "cd /tmp/tune/fio && atune-adm tuning --project fio --detail tuning_fio_client.yaml",
			Restore:       "cd /tmp/tune/fio && atune-adm tuning --restore --project fio",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [test_path] [test_disk]",
	}
	return info
}
