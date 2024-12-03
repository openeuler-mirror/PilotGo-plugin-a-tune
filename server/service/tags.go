/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
package service

import (
	"gitee.com/openeuler/PilotGo/sdk/common"
	"openeuler.org/PilotGo/atune-plugin/dao"
	"openeuler.org/PilotGo/atune-plugin/plugin"
)

func GetTags() {
	tag_cb := func(uuids []string) []common.Tag {
		var tags []common.Tag
		for _, uuid := range uuids {
			ok, _ := dao.IsExist(uuid)
			if ok {
				tag := common.Tag{
					UUID: uuid,
					Type: common.TypeOk,
					Data: "atune",
				}
				tags = append(tags, tag)
			} else {
				tag := common.Tag{
					UUID: uuid,
					Type: common.TypeError,
					Data: "",
				}
				tags = append(tags, tag)
			}
		}
		return tags
	}
	plugin.GlobalClient.OnGetTags(tag_cb)
}
