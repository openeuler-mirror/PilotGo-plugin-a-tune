package dao

import (
	"time"

	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func SaveRusult(result *model.RunResult) error {
	if err := db.MySQL().Create(&result).Error; err != nil {
		return err
	}
	return nil
}
func UpdateTaskResultCount(taskId int) error {
	var successCount int64

	if err := db.MySQL().Model(&model.RunResult{}).Where("task_id = ? AND command_type = ? AND is_success = ?", taskId, "restore", "success").Count(&successCount).Error; err != nil {
		return err
	}
	task := &model.Tasks{
		UpdateTime:   time.Now().Format("2006-01-02 15:04:05"),
		SuccessCount: int(successCount),
	}
	if err := db.MySQL().Model(&model.Tasks{}).Where("id = ?", taskId).Updates(&task).Error; err != nil {
		return err
	}
	return nil
}

func UpdateResult(dbtaskid int, machine_uuid string, commandType string, result *model.RunResult) error {
	var r model.RunResult
	if err := db.MySQL().Model(&r).Where("task_id = ? AND machine_uuid = ? AND command_type = ?", dbtaskid, machine_uuid, commandType).Updates(&result).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResult(resultId int) error {
	var r model.RunResult
	if err := db.MySQL().Where("id = ?", resultId).Unscoped().Delete(&r).Error; err != nil {
		return err
	}
	return nil
}

func SearchResult(taskId string) ([]*model.RunResult, error) {
	var result []*model.RunResult
	if err := db.MySQL().Where("task_id = ?", taskId).Find(&result).Error; err != nil {
		return []*model.RunResult{}, nil
	}

	return result, nil
}
