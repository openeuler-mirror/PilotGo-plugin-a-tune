/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package dao

import (
	"errors"

	"openeuler.org/PilotGo/atune-plugin/db"
	"openeuler.org/PilotGo/atune-plugin/model"
)

func IsExist(uuid string) (bool, error) {
	var ac model.AtuneClient
	err := db.MySQL().Where("machine_uuid = ?", uuid).Find(&ac).Error
	if err != nil {
		return false, errors.New("查询数据库失败：" + err.Error())
	}
	return ac.ID != 0, nil
}

func AddAtuneClientList(ac *model.AtuneClient) error {
	a := model.AtuneClient{
		MachineUUID: ac.MachineUUID,
	}
	err := db.MySQL().Save(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAtuneClientList(ac *model.AtuneClient) error {
	var a model.AtuneClient
	err := db.MySQL().Where("machine_uuid = ?", ac.MachineUUID).Unscoped().Delete(a).Error
	if err != nil {
		return err
	}
	return nil
}
