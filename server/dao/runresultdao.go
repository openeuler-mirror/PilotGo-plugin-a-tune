/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
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

func SearchResult(taskId, machine_uuid string) ([]*model.RunResult, error) {
	var result []*model.RunResult
	if err := db.MySQL().Where("task_id = ? AND machine_uuid = ?", taskId, machine_uuid).Find(&result).Error; err != nil {
		return []*model.RunResult{}, nil
	}

	return result, nil
}

func GetResultUUIDByTaskId(taskId string) ([]string, error) {
	var uuids []string
	err := db.MySQL().Model(&model.RunResult{}).Distinct("machine_uuid").Where("task_id = ?", taskId).Pluck("machine_uuid", &uuids).Error
	return uuids, err
}

func GetResultByTaskIdAndUUID(taskId int, machine_uuid string, commandType string) (string, error) {
	var r model.RunResult
	err := db.MySQL().Where("task_id = ? AND machine_uuid = ? AND command_type = ?", taskId, machine_uuid, commandType).Find(&r).Error
	return r.Stdout, err
}
