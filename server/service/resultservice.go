package service

import (
	"errors"
	"strconv"

	"gitee.com/openeuler/PilotGo/sdk/common"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"gitee.com/openeuler/PilotGo/sdk/response"
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

func UpdateResultStatus(dbtaskid int, uuid string, commandType string) error {
	result := &model.RunResult{
		IsSuccess: IsSuccess_running,
	}

	err := dao.UpdateResult(dbtaskid, uuid, commandType, result)
	if err != nil {
		return errors.New("更新执行任务状态失败：" + err.Error())
	}
	return nil
}
func SearchResult(searchKey string, query *response.PaginationQ) ([]*model.RunResult, int, error) {
	if data, total, err := dao.SearchResult(searchKey, query); err != nil {
		return nil, 0, err
	} else {
		return data, int(total), nil
	}
}
