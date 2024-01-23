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
