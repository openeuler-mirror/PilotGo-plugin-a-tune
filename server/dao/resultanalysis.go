package dao

import (
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func SaveResultAnalysis(ra *model.TuningResults) error {
	if err := db.MySQL().Create(&ra).Error; err != nil {
		return err
	}
	return nil
}

func UpdateResultAnalysis(taskId int, machine_uuid string, ra *model.TuningResults) error {
	var r model.TuningResults
	if err := db.MySQL().Model(&r).Where("task_id = ? AND machine_uuid = ?", taskId, machine_uuid).Updates(&ra).Error; err != nil {
		return err
	}
	return nil
}

func IsExistResultAnalisis(taskId int, machine_uuid string) (bool, error) {
	var result model.TuningResults
	err := db.MySQL().Where("task_id = ? AND machine_uuid = ?", taskId, machine_uuid).Find(&result).Error
	return result.ID != 0, err
}

func GetResultAnalisis(taskId, machine_uuid string) (model.TuningResults, error) {
	var result model.TuningResults
	if err := db.MySQL().Where("task_id = ? AND machine_uuid = ?", taskId, machine_uuid).Find(&result).Error; err != nil {
		return model.TuningResults{}, nil
	}
	return result, nil
}
