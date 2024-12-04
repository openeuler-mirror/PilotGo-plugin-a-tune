/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
import { defineStore } from "pinia";
import { Task, Atune } from "@/types/atune";

export const useRouterStore = defineStore("router", {
  state: () => ({
    route: "",
  }),
  getters: {
    currentRoute: (state) => state.route,
  },
  actions: {
    setCurrentRoute(route: string) {
      this.route = route;
    },
    showRoute(path: string, str: string) {
      return path.indexOf(str) >= 0;
    },
  },
});
