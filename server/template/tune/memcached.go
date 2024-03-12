package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type MemcachedApp struct{}

func (m *MemcachedApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "memcached",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e memcached.tar.gz ] && tar -xzvf memcached.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/memcached.tar.gz && tar -xzvf memcached.tar.gz)",
		Prepare:       "cd /tmp/tune/memcached && sh prepare.sh",
		Tune:          "cd /tmp/tune/memcached && atune-adm tuning --project memcached_memaslap --detail memcached_memaslap_client.yaml",
		Restore:       "cd /tmp/tune/memcached && atune-adm tuning --restore --project memcached_memaslap",
	}
	return info
}
