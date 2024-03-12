package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type TomcatApp struct{}

type TomcatImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (tomcat *TomcatApp) Info() *TomcatImp {
	info := &TomcatImp{
		BaseTune: TuneInfo{
			TuneName:      "tomcat",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e tomcat.tar.gz ] && tar -xzvf tomcat.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/tomcat.tar.gz && tar -xzvf tomcat.tar.gz)",
			Prepare:       "cd /tmp/tune/tomcat && sh prepare.sh tomcat_root",
			Tune:          "cd /tmp/tune/tomcat && atune-adm tuning --project tomcat --detail tomcat.yaml",
			Restore:       "cd /tmp/tune/tomcat && atune-adm tuning --restore --project tomcat",
		},
		Notes: "注意prepare.sh的使用方法, sh prepare.sh [tomcat 目录名称], 默认是tomcat_root",
	}
	return info
}
