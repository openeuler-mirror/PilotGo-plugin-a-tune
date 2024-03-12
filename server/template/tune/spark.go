package tune

import "openeuler.org/PilotGo/atune-plugin/plugin"

type SparkApp struct{}

type SparkImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (spark *SparkApp) Info() *SparkImp {
	info := &SparkImp{
		BaseTune: TuneInfo{
			TuneName:      "spark",
			WorkDirectory: "mkdir -p /tmp/tune/ && cd /tmp/tune/ && [ -e spark.tar.gz ] && tar -xzvf spark.tar.gz || ( curl -OJ http://" + plugin.GlobalClient.Server() + "/api/v1/download/spark.tar.gz && tar -xzvf spark.tar.gz)",
			Prepare:       "cd /tmp/tune/spark && sh run_env.sh",
			Tune:          "",
			Restore:       "",
		},
		Notes: "调优步骤均在run_env.sh完成, 执行该脚本可完成环境的部署和调优测试",
	}
	return info
}
