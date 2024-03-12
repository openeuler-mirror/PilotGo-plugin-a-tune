package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type TidbApp struct{}

type TidbImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (tidb *TidbApp) Info() *TidbImp {
	info := &TidbImp{
		BaseTune: TuneInfo{
			TuneName:      "tidb",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e tidb.tar.gz ] && tar -xzvf tidb.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/tidb.tar.gz && tar -xzvf tidb.tar.gz)",
			Prepare:       "cd /tmp/tune/tidb && sh prepare.sh",
			Tune:          "cd /tmp/tune/tidb && atune-adm tuning --project tidb --detail tidb_client.yaml",
			Restore:       "cd /tmp/tune/tidb && atune-adm tuning --restore --project tidb",
		},
		Notes: "如果目标机器是第一次使用tidb调优, 请为目标机器准备好测试环境, 详情请查看 https://gitee.com/openeuler/A-Tune/blob/master/examples/tuning/tidb/README",
	}
	return info
}
