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
	MachineIP   string `json:"machine_ip"`
	CommandType string `json:"command_type"`
	RetCode     int    `json:"retcode"`
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	IsSuccess   string `json:"is_success"`
}

type Tasks struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TuneID     int    `json:"tune_id"`
	TaskName   string `json:"task_name"`
	TaskStatus string `json:"task_status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	// RunResults []RunResult `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE;" json:"results"`
	Tune         Tunes `gorm:"foreignKey:TuneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"tune"`
	SuccessCount int   `json:"success_count"`
	UUIDCount    int   `json:"uuid_count"`
}

type AtuneClient struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MachineUUID string `json:"machine_uuid"`
	MachineIP   string `json:"machine_ip"`
}

type TaskCommand struct {
	PrepareCommand string `json:"prepare"`
	TuneCommand    string `json:"tune"`
	RestoreCommand string `json:"restore"`
}