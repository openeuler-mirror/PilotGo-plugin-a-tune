/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package dao

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

// 根据id获取调优命令
func GetTuneComandById(tune_id int) (map[int]model.TaskCommand, error) {
	var t model.Tunes
	if err := db.MySQL().Where("id = ?", tune_id).Find(&t).Error; err != nil {
		return nil, err
	}
	tuneCommand := make(map[int]model.TaskCommand)
	tuneCommand[tune_id] = model.TaskCommand{
		PrepareCommand: t.WorkDirectory + " && " + t.Prepare,
		TuneCommand:    t.Tune,
		RestoreCommand: t.Restore,
	}
	return tuneCommand, nil
}

func QueryTunesNoPaged() ([]*model.Tunes, error) {
	var tunes []*model.Tunes
	if err := db.MySQL().Find(&tunes).Error; err != nil {
		return []*model.Tunes{}, err
	}
	return tunes, nil
}

func QueryTunes(query *response.PaginationQ) ([]*model.Tunes, int64, error) {
	var tunes []*model.Tunes
	if err := db.MySQL().Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&tunes).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.MySQL().Model(&tunes).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return tunes, total, nil
}

func SaveTune(tune *model.Tunes) error {
	if err := db.MySQL().Create(&tune).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTune(tuneName string, tune *model.Tunes) error {
	var t model.Tunes
	if err := db.MySQL().Model(&t).Where("tune_name = ?", tuneName).Updates(&tune).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTune(tuneId int) error {
	var t model.Tunes
	if err := db.MySQL().Where("id = ?", tuneId).Unscoped().Delete(&t).Error; err != nil {
		return err
	}
	return nil
}

func SearchTune(search string, query *response.PaginationQ) ([]*model.Tunes, int64, error) {
	var tune []*model.Tunes
	if err := db.MySQL().Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("tune_name LIKE ? OR description LIKE ? OR custom_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&tune).Error; err != nil {
		return nil, 0, nil
	}

	var total int64
	if err := db.MySQL().Where("tune_name LIKE ? OR description LIKE ? OR custom_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Model(&tune).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return tune, total, nil
}

func QueryTuneMachines() ([]*model.AtuneClient, error) {
	var ac []*model.AtuneClient
	if err := db.MySQL().Find(&ac).Error; err != nil {
		return []*model.AtuneClient{}, err
	}
	return ac, nil
}
