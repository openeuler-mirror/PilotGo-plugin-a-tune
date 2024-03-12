package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type GoGcApp struct{}

func (g *GoGcApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "go_gc",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e go_gc.tar.gz ] && tar -xzvf go_gc.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/go_gc.tar.gz && tar -xzvf go_gc.tar.gz)",
		Prepare:       "cd /tmp/tune/go_gc && sh prepare.sh",
		Tune:          "cd /tmp/tune/go_gc && atune-adm tuning --project go_gc --detail go_gc_client.yaml",
		Restore:       "cd /tmp/tune/go_gc && atune-adm tuning --restore --project go_gc",
	}
	return info
}
