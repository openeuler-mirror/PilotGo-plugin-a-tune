/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
import { defineStore } from "pinia";
import { Task } from "@/types/atune";

export const useAtuneStore = defineStore("atune", {
  state: () => ({
    taskRow: {} as Task,
  }),
  actions: {
    // 设置taskRow数据
    setTaskRow(row: Task) {
      this.taskRow = row;
    },

  },
  persist: true
});
