package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type MariadbApp struct{}

type MariadbImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (m *MariadbApp) Info() *MariadbImp {
	info := &MariadbImp{
		BaseTune: TuneInfo{
			TuneName:      "mariadb",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e mariadb.tar.gz ] && tar -xzvf mariadb.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/mariadb.tar.gz && tar -xzvf mariadb.tar.gz)",
			Prepare:       "cd /tmp/tune/mariadb && sh prepare.sh 25",
			Tune:          "cd /tmp/tune/mariadb && atune-adm tuning --project mariadb --detail mariadb_client.yaml",
			Restore:       "cd /tmp/tune/mariadb && atune-adm tuning --restore --project mariadb",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [迭代次数, 默认25]",
	}
	return info
}
