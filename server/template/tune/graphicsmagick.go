package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type GraphicsmagickApp struct{}

func (g *GraphicsmagickApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "graphicsmagick",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e graphicsmagick.tar.gz ] && tar -xzvf graphicsmagick.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/graphicsmagick.tar.gz && tar -xzvf graphicsmagick.tar.gz)",
		Prepare:       "cd /tmp/tune/graphicsmagick && sh prepare.sh",
		Tune:          "cd /tmp/tune/graphicsmagick && atune-adm tuning --project graphicsmagick  --detail gm_client.yaml",
		Restore:       "cd /tmp/tune/graphicsmagick && atune-adm tuning --restore --project graphicsmagick",
	}
	return info
}
