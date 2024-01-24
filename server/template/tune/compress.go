package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type CompressApp struct{}

func (c *CompressApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "compress",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e compress.tar.gz ] && tar -xzvf compress.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/compress.tar.gz && tar -xzvf compress.tar.gz)",
		Prepare:       "cd /tmp/tune/compress && sh prepare.sh enwik8.zip",
		Tune:          "cd /tmp/tune/compress && atune-adm tuning --project compress --detail compress_client.yaml",
		Restore:       "cd /tmp/tune/compress && atune-adm tuning --restore --project compress",
	}
	return info
}
