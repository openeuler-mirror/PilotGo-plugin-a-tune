package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type NginxApp struct{}

type NginxImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (nginx *NginxApp) Info() *NginxImp {
	info := &NginxImp{
		BaseTune: TuneInfo{
			TuneName:      "nginx",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e nginx.tar.gz ] && tar -xzvf nginx.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/nginx.tar.gz && tar -xzvf nginx.tar.gz)",
			Prepare:       "cd /tmp/tune/nginx && sh prepare.sh",
			Tune:          "cd /tmp/tune/nginx && atune-adm tuning --project nginx --detail nginx_client.yaml",
			Restore:       "cd /tmp/tune/nginx && atune-adm tuning --restore --project nginx",
		},
		Notes: "nginx长连接:【调优指令】atune-adm tuning --project nginx_http_long --detail nginx_http_long_client.yaml【恢复环境】atune-adm tuning --restore --project nginx_http_long",
	}
	return info
}
