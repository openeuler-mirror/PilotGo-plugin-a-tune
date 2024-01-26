package model

type Tunes struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TuneName      string `json:"tuneName"`
	CustomName    string `json:"custom_name"`
	Description   string `json:"description"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
	WorkDirectory string `json:"workDir"`
	Prepare       string `json:"prepare"`
	Tune          string `json:"tune"`
	Restore       string `json:"restore"`
	Notes         string `json:"note"`
}

type RunResult struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TaskID      int    `json:"task_id"`
	MachineUUID string `json:"machine_uuid"`
	CommandType string `json:"command_type"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	RetCode     int    `json:"retcode"`
	Stdout      string `json:"stdout"`
	IsSuccess   string `json:"is_success"`
}

type Tasks struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TuneID      int    `json:"tune_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	TaskStatus  string `json:"task_status"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	// RunResults []RunResult `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;" json:"results"`
	Tune         Tunes `gorm:"foreignKey:TuneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"tune"`
	SuccessCount int   `json:"success_count"`
	UUIDCount    int   `json:"uuid_count"`
}

type AtuneClient struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MachineUUID string `json:"machine_uuid"`
}

type TaskCommand struct {
	PrepareCommand string `json:"prepare"`
	TuneCommand    string `json:"tune"`
	RestoreCommand string `json:"restore"`
}

type Results struct {
	TaskID      string       `json:"taskId"`
	MachineUUID string       `json:"machineUUID"`
	MachineIP   string       `json:"machineIp"`
	CPUArch     string       `json:"cpuArch"`
	OS          string       `json:"os"`
	Result      []*RunResult `json:"result"`
}

// TuningIteration 是用于存储每次迭代结果的结构体
type TuningIterResult struct {
	TuningStep             string      `json:"step"`
	RecommandParameters    interface{} `json:"recommand_parameters"`
	EvaluationValue        interface{} `json:"evaluation_value"`
	PerformanceImprovement string      `json:"performance_improvement"`
}

// TuningResult 是用于存储调优结果的结构体
type TuningResults struct {
	ID                      int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TaskID                  int    `json:"task_id"`
	MachineUUID             string `json:"machine_uuid"`
	TuningResult            string `json:"tuning_result"`
	FinalOptimizationResult string `json:"final_optimization_result"`
	FinalEvaluationValue    string `json:"final_evaluation_value"`
	BaselinePerformance     string `json:"baseline_performance"`
}
