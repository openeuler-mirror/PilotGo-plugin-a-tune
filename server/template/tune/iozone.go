package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type IozoneApp struct{}

type IozoneImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (io *IozoneApp) Info() *IozoneImp {
	info := &IozoneImp{
		BaseTune: TuneInfo{
			TuneName:      "iozone",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e iozone.tar.gz ] && tar -xzvf iozone.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/iozone.tar.gz && tar -xzvf iozone.tar.gz)",
			Prepare:       "cd /tmp/tune/fio && sh prepare.sh 参数一 参数二",
			Tune:          "cd /tmp/tune/fio && atune-adm tuning --project iozone --detail tuning_iozone_client.yaml",
			Restore:       "cd /tmp/tune/fio && atune-adm tuning --restore --project iozone",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [test_path] [diskname]",
	}
	return info
}
