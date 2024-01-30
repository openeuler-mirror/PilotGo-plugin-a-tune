package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type CompressExceptApp struct{}

func (ce *CompressExceptApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "compress_Except_example",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e compress_Except_example.tar.gz ] && tar -xzvf compress_Except_example.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/compress_Except_example.tar.gz  && tar -xzvf compress_Except_example.tar.gz)",
		Prepare:       "cd /tmp/tune/compress_Except_example && sh prepare.sh enwik8.zip",
		Tune:          "cd /tmp/tune/compress_Except_example && atune-adm tuning --project compress_Except_example --detail compress_Except_example_client.yaml",
		Restore:       "cd /tmp/tune/compress_Except_example && atune-adm tuning --restore --project compress_Except_example",
	}
	return info

}
