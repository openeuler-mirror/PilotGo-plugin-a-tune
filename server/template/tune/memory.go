package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type MemoryApp struct{}

func (m *MemoryApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "memory",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e memory.tar.gz ] && tar -xzvf memory.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/memory.tar.gz && tar -xzvf memory.tar.gz)",
		Prepare:       "cd /tmp/tune/memory && sh prepare.sh",
		Tune:          "cd /tmp/tune/memory && atune-adm tuning --project stream --detail tuning_stream_client.yaml",
		Restore:       "cd /tmp/tune/memory && atune-adm tuning --restore --project stream",
	}
	return info
}
