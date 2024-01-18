package service

import (
	"errors"
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/model"
)

const (
	// 远程执行命令状态
	IsSuccess_waiting = "waiting"
	IsSuccess_running = "running"
	IsSuccess_fail    = "fail"
	IsSuccess_success = "success"
)

const (
	// 命令类型
	CommandTypePrepare = "prepare"
	CommandTypeTune    = "tune"
	CommandTypeRestore = "restore"
)

func DeleteResult(resultId []int) error {
	if len(resultId) == 0 {
		return errors.New("请输入删除的ID")
	}

	for _, rid := range resultId {
		if err := dao.DeleteResult(rid); err != nil {
			logger.Error("%v", strconv.Itoa(rid)+"未删除成功")
		}
	}
	return nil
}

func processResult(dbtaskid int, res *common.CmdResult, commandType string) (string, error) {
	var commandResultStatus string
	if res.Stderr == "" && len(res.Stdout) != 0 {
		commandResultStatus = IsSuccess_success
	} else {
		commandResultStatus = IsSuccess_fail
	}
	result := &model.RunResult{
		MachineIP: res.MachineIP,
		RetCode:   res.RetCode,
		Stdout:    res.Stdout,
		Stderr:    res.Stderr,
		IsSuccess: commandResultStatus,
	}

	if err := dao.UpdateResult(dbtaskid, res.MachineUUID, commandType, result); err != nil {
		return commandResultStatus, errors.New("更新命令执行结果失败：" + err.Error())
	}

	return commandResultStatus, nil
}

func UpdateResultStatusToRunning(dbtaskid int, uuid string, commandType string) error {
	result := &model.RunResult{
		IsSuccess: IsSuccess_running,
	}

	err := dao.UpdateResult(dbtaskid, uuid, commandType, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	return nil
}

func UpdateResultStatusForPrepare(dbtaskid int, uuid string, commandType string, resultStatus string) error {
	result := &model.RunResult{
		IsSuccess: resultStatus,
	}

	err := dao.UpdateResult(dbtaskid, uuid, commandType, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}

	err = dao.UpdateResult(dbtaskid, uuid, CommandTypeTune, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	err = dao.UpdateResult(dbtaskid, uuid, CommandTypeRestore, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	return nil
}
func UpdateResultStatusForTune(dbtaskid int, uuid string, commandType string, resultStatus string) error {
	result := &model.RunResult{
		IsSuccess: resultStatus,
	}

	err := dao.UpdateResult(dbtaskid, uuid, commandType, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	err = dao.UpdateResult(dbtaskid, uuid, CommandTypeRestore, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	return nil
}

func UpdateResultStatus(dbtaskid int, uuid string, commandType string, resultStatus string) error {
	result := &model.RunResult{
		IsSuccess: resultStatus,
	}

	err := dao.UpdateResult(dbtaskid, uuid, commandType, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	if commandType == CommandTypeRestore {
		if err := dao.UpdateTaskResultCount(dbtaskid); err != nil {
			return err
		}
	}
	return nil
}
func SearchResultByTaskId(taskId string) ([]*model.RunResult, error) {
	if data, err := dao.SearchResult(taskId); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
