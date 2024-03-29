package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type GccComplieApp struct{}

func (gcc *GccComplieApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "gcc_compile",
		WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e gcc_compile.tar.gz ] && tar -xzvf gcc_compile.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/gcc_compile.tar.gz && tar -xzvf gcc_compile.tar.gz)",
		Prepare:       "cd /tmp/tune/gcc_compile && sh prepare.sh",
		Tune:          "cd /tmp/tune/gcc_compile && atune-adm tuning --project gcc_compile --detail gcc_compile_client.yaml",
		Restore:       "cd /tmp/tune/gcc_compile && atune-adm tuning --restore --project gcc_compile",
	}
	return info
}
func (gcc *GccComplieApp) Prepare() error {
	return nil
}
func (gcc *GccComplieApp) StartTune() error {
	return nil
}
func (gcc *GccComplieApp) Restore() error {
	return nil
}
