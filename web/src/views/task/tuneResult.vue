<!--
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhaozhenfang <zhaozhenfang@kylinos.cn>
 * Date: Tue Jan 30 14:01:39 2024 +0800
-->
<template>
  <div style="width:98%; height: 100%; overflow: hidden;">
    <el-descriptions :column="1" style="padding-left: 10px;">
      <el-descriptions-item label="最优参数：">
        <span v-for="(key, index) in Object.keys(final_optimization)">
          {{ key + ':' + final_optimization[key] }}
          <el-divider direction="vertical"></el-divider>
        </span>
      </el-descriptions-item>
      <el-descriptions-item label="基准性能：">
        <span v-for="(key, index) in Object.keys(baseline)">
          {{ key + ':' + baseline[key] }}
          <el-divider direction="vertical"></el-divider>
        </span></el-descriptions-item>
      <el-descriptions-item label="最终评估值："><span v-for="(key, index) in Object.keys(final_evaluation)">
          {{ key + ':' + final_evaluation[key] }}
          <el-divider direction="vertical"></el-divider>
        </span></el-descriptions-item>
    </el-descriptions>
    <el-table :data="tableData" height="90%" style="width: 100%">
      <el-table-column prop="step" label="迭代轮次" width="180" />
      <el-table-column label="调优参数">
        <template #default="{ row }">
          <span v-for="key in Object.keys(row.recommand_parameters)">
            {{ key + ':' + row.recommand_parameters[key] }}
            <el-divider direction="vertical"></el-divider>
          </span>
        </template>
      </el-table-column>
      <el-table-column label="评估值">
        <template #default="{ row }">
          <span v-for="key in Object.keys(row.evaluation_value)">
            {{ key + ':' + row.evaluation_value[key] }}
            <el-divider direction="vertical"></el-divider>
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="performance_improvement" sortable label="性能提升率" width="120" :sort-method="(a: any, b: any) => {
        return a.performance_improvement - b.performance_improvement;
      }">
        <template #default="props">
          <span class="flex" style="justify-content: space-between;">
            {{ props.row.performance_improvement + '%' }}
            <el-icon color="#67C23A" v-if="props.row.performance_improvement > 0">
              <Top />
            </el-icon>
            <el-icon color="#F56C6C" v-else>
              <Bottom />
            </el-icon></span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { ElMessage } from "element-plus";
import { Top, Bottom } from '@element-plus/icons-vue'
import { getTuneResult } from "@/api/result";

const props = defineProps({
  taskId: {
    type: Number,
    default: null,
    required: true
  },
  tune_uuid: {
    type: String,
    default: '',
    required: true
  }
})

const baseline: any = ref({});
const final_evaluation: any = ref({});
const final_optimization: any = ref({});

onMounted(() => {
  getFinalResult();
})
const tableData = ref([] as any)
// 获取机器的最终调优结果
const getFinalResult = () => {
  getTuneResult({ taskId: props.taskId, uuid: props.tune_uuid }).then((res: any) => {
    if (res.data.code === 200) {
      let { baseline_performance, final_evaluation_value, final_optimization_result, tuning_result } = res.data.data;
      baseline.value = baseline_performance;
      final_evaluation.value = final_evaluation_value;
      final_optimization.value = final_optimization_result;
      tableData.value = tuning_result;
    } else {
      ElMessage.error(res.data.msg);
    }
  })
}
</script>

<style scoped></style>