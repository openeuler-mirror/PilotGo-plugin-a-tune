package service

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/model"
	"openeuler.org/PilotGo/atune-plugin/plugin"
)

const (
	// 任务状态
	TaskWaiting = "waiting"
	TaskRunnig  = "running"
	TaskEnd     = "ending"
)

func RunTask(dbtaskId int, tuneId int, commands map[int]model.TaskCommand, machine_uuids []string) error {
	if len(machine_uuids) == 0 {
		return errors.New("机器uuid未选中")
	}

	// 使用 channel 来通知任务完成
	taskComplete := make(chan struct{})
	// 定义成功和失败任务计数器
	var successCount, failCount int
	var mu sync.Mutex
	for _, uuid := range machine_uuids {
		go func(uuid string) {
			defer func() {
				taskComplete <- struct{}{}
			}()

			isSuccess_prepare, err := runCommandAndProcessResult(dbtaskId, uuid, commands[tuneId].PrepareCommand, CommandTypePrepare)
			if err != nil || isSuccess_prepare == IsSuccess_fail {
				mu.Lock()
				failCount++
				mu.Unlock()
				return
			}
			isSuccess_tune, err := runCommandAndProcessResult(dbtaskId, uuid, commands[tuneId].TuneCommand, CommandTypeTune)
			if err != nil || isSuccess_tune == IsSuccess_fail {
				mu.Lock()
				failCount++
				mu.Unlock()
				return
			}
			isSuccess_restore, err := runCommandAndProcessResult(dbtaskId, uuid, commands[tuneId].RestoreCommand, CommandTypeRestore)
			if err != nil || isSuccess_restore == IsSuccess_fail {
				mu.Lock()
				failCount++
				mu.Unlock()
				return
			}
			mu.Lock()
			successCount++
			mu.Unlock()
		}(uuid)
	}

	go func() {
		var completedTasks int
		for range machine_uuids {
			<-taskComplete
			completedTasks++

			// 判断是否任务都已完成
			if completedTasks == len(machine_uuids) {
				UpdateTaskStatus(dbtaskId, TaskEnd)
			}
		}
	}()
	return nil
}

func runCommandAndProcessResult(dbtaskId int, machine_uuid string, cmd string, commandType string) (string, error) {
	b := &common.Batch{
		MachineUUIDs: []string{machine_uuid},
	}
	err := UpdateResultStatus(dbtaskId, machine_uuid, commandType)
	if err != nil {
		logger.Error("%v", err.Error())
		return "", err
	}
	res, err := plugin.GlobalClient.RunCommand(b, cmd)
	if err != nil {
		logger.Error("%v远程执行命令%v失败: %v", machine_uuid, cmd, err.Error())
		return "", err
	}

	isSuccess, err := processResult(dbtaskId, res[0], commandType)
	if err != nil {
		logger.Error("%v执行结果保存到数据库失败: %v", machine_uuid, err.Error())
		return "", err
	}
	return isSuccess, nil
}

func QueryTaskLists(query *response.PaginationQ) ([]*model.Tasks, int, error) {
	if data, total, err := dao.QueryTaskLists(query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}

func UpdateTaskStatus(taskId int, taskStatus string) error {
	task := &model.Tasks{
		TaskStatus: taskStatus,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := dao.UpdateTask(taskId, task)
	if err != nil {
		return errors.New("保存执行任务状态失败：" + err.Error())
	}
	return nil
}
func SaveTask(task_name string, uuids []string, tuneId int) error {
	task := &model.Tasks{
		TaskName:   task_name,
		TuneID:     tuneId,
		TaskStatus: TaskWaiting,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	taskid, err := dao.SaveTask(task)
	if err != nil {
		return errors.New("保存执行任务失败：" + err.Error())
	}

	for _, uuid := range uuids {
		result_prepare := &model.RunResult{
			TaskID:      taskid,
			MachineUUID: uuid,
			CommandType: CommandTypePrepare,
			IsSuccess:   IsSuccess_waiting,
		}
		if err := dao.SaveRusult(result_prepare); err != nil {
			logger.Error("save result uuid failed: %v", err.Error())
		}
		result_tune := &model.RunResult{
			TaskID:      taskid,
			MachineUUID: uuid,
			CommandType: CommandTypeTune,
			IsSuccess:   IsSuccess_waiting,
		}
		if err := dao.SaveRusult(result_tune); err != nil {
			logger.Error("save result uuid failed: %v", err.Error())
		}
		result_restore := &model.RunResult{
			TaskID:      taskid,
			MachineUUID: uuid,
			CommandType: CommandTypeRestore,
			IsSuccess:   IsSuccess_waiting,
		}
		if err := dao.SaveRusult(result_restore); err != nil {
			logger.Error("save result uuid failed: %v", err.Error())
		}
	}
	return nil
}
func DeleteTask(taskId []int) error {
	if len(taskId) == 0 {
		return errors.New("请输入调优模板ID")
	}

	for _, tid := range taskId {
		if err := dao.DeleteTask(tid); err != nil {
			logger.Error("%v", strconv.Itoa(tid)+"未删除成功")
		}
	}
	return nil
}

func SearchTask(search string, query *response.PaginationQ) ([]*model.Tasks, int, error) {
	if data, total, err := dao.SearchTask(search, query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}
