/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
export interface Task {
  id: number;
  uuid_count: number;
  success_count: number;
  task_name: string;
  command: string;
  task_status: string;
  create_time: string;
  update_time: string;
  results: Array;
  tune_id: int;
  tune: Atune;
  machine_uuids: Array[string];
  description: string;
}

export interface Atune {
  id: number;
  description: string;
  create_time: string;
  update_time: string;
  tuneName: string;
  custom_name: string;
  workDir: string;
  prepare: string;
  tune: string;
  restore: string;
  note: string;
}

// 画布矩形的参数配置
export interface RectConfig {
  x: number;
  y: number;
  width: number;
  height: number;
  fill: string;
  stroke: string;
  shadowBlur: number;
  cornerRadius: number;
}

// 导出接口字面量
export type TaskArray = Task[];
export type AtuneArray = Atune[];

// *接口api返回结果约束不含data
export interface Result {
  code: number;
  msg?: string;
}

// *接口api返回结果含有page信息
export interface ResultData extends Result {
  data: Array;
  ok?: Boolean;
  page: number;
  size: number;
  total: number;
}
