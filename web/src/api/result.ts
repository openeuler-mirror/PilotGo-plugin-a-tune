/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
import request from './request';

// 获取任务结果详情
export function getTaskDetail(data: object) {
  return request({
    url: '/plugin/atune/result',
    method: 'get',
    params: data,
  });
}

// 获取调优结果详情
export function getTuneResult(data: object) {
  return request({
    url: '/plugin/atune/tune_result',
    method: 'get',
    params: data,
  });
}