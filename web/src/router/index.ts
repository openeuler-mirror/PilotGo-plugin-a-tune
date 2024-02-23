import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import Home from "@/views/Home.vue";
import Atune from "@/views/atuneList.vue";
import Result from "@/views/ResultInfo.vue";

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
    meta: { title: "模板列表" },
    children: [
      {
        path: "detail",
        name: "atuneDetail",
        component: () => import("../views/atuneDetail.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(window.__MICRO_APP_BASE_ROUTE__ || '/'),
  routes,
});

export default router;
