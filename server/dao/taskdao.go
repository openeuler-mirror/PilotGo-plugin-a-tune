package dao

import (
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func QueryTaskLists(query *response.PaginationQ) ([]*model.Tasks, int64, error) {
	var tasks []*model.Tasks
	if err := db.MySQL().Preload("Tune").Order("id desc").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.MySQL().Model(&tasks).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}

func SaveTask(task *model.Tasks) (int, error) {
	if err := db.MySQL().Create(&task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil
}
func UpdateTask(dbtaskid int, task *model.Tasks) error {
	var t model.Tasks
	if err := db.MySQL().Model(&t).Where("id = ?", dbtaskid).Updates(&task).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(taskId int) error {
	var task model.Tasks

	// 手动删除关联的 RunResults 数据
	if err := db.MySQL().Unscoped().Where("task_id = ?", taskId).Delete(&model.RunResult{}).Error; err != nil {
		return err
	}

	if err := db.MySQL().Unscoped().Where("id = ?", taskId).Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func SearchTask(search string, query *response.PaginationQ) ([]*model.Tasks, int64, error) {
	var task []*model.Tasks
	if err := db.MySQL().Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("task_name LIKE ? OR script LIKE ? OR task_status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&task).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.MySQL().Where("task_name LIKE ? OR script LIKE ? OR task_status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Model(&task).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return task, total, nil
}
