/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package tune

var TuneGroupApp = new(TuneGroup)

type TuneGroup struct {
	Compress                   CompressApp
	CompressExcept             CompressExceptApp
	Ffmpeg                     FfmpegApp
	Fio                        FioApp
	GccCompile                 GccComplieApp
	GoGc                       GoGcApp
	Graphicsmagick             GraphicsmagickApp
	Iozone                     IozoneApp
	Kafka                      KafkaApp
	KeyParametersSelect        KeyParametersSelectApp
	KeyParametersSelectVariant KeyParametersSelectVariantApp
	Mariadb                    MariadbApp
	Memcached                  MemcachedApp
	Memory                     MemoryApp
	MysqlSysbench              MysqlSysbenchApp
	Nginx                      NginxApp
	OpenGauss                  OpenGaussApp
	Redis                      RedisApp
	Spark                      SparkApp
	TensorflowTrain            TensorflowTrainApp
	Tidb                       TidbApp
	Tomcat                     TomcatApp
}

type TuneInfo struct {
	TuneName      string `json:"tuneName"`
	Description   string `json:"description"`
	WorkDirectory string `json:"workDir"`
	Prepare       string `json:"prepare"`
	Tune          string `json:"tune"`
	Restore       string `json:"restore"`
}
