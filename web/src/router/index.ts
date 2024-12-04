/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import Home from "@/views/Home.vue";
import Atune from "@/views/atuneList.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "",
    redirect: "/task",
  },
  {
    path: "/task",
    component: Home,
    meta: { title: "任务列表" },
    children: [
      {
        path: "detail",
        name: "taskDetail",
        component: () => import("../views/taskDetail.vue"),
      },
    ],
  },
  {
    path: "/template",
    component: Atune,
    meta: { title: "模板列表" }
  },
];
const router = createRouter({
  history: createWebHistory(window.__MICRO_APP_BASE_ROUTE__ || '/'),
  routes,
});

export default router;
