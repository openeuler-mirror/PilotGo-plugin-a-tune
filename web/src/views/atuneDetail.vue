<template>
  <div class="atuneDetail">
    <el-container class="atuneDetail_con">
      <el-header height="20%">
        <el-descriptions :column="1">
          <el-descriptions-item label="模板名称：">
            {{ custom_name }}</el-descriptions-item>
          <el-descriptions-item label="调优对象：">
            {{ tuneName }}</el-descriptions-item>
          <el-descriptions-item label="工作目录：">{{
            workDir
          }}</el-descriptions-item>
          <el-descriptions-item label="注意事项：">{{
            note
          }}</el-descriptions-item>
          <el-descriptions-item>
            <el-divider />
          </el-descriptions-item>
          <el-descriptions-item label="流程如下：">
            <el-steps align-center :active="active" finish-status="success">
              <el-step v-for="(item, index) in stepArr" :key="index" :title="item" />
            </el-steps>
          </el-descriptions-item>
          <el-descriptions-item :label="stepArr[active - 1] + '：'">
            <el-tag type="" effect="dark" size="large">
              {{ shell }}
            </el-tag></el-descriptions-item>
          <el-descriptions-item>
            <el-divider content-position="center">
              <el-button type="primary" plain size="small" style="margin-top: 12px" @click="next">下一步</el-button>
            </el-divider>
          </el-descriptions-item>
        </el-descriptions>
      </el-header>
    </el-container>
  </div>
</template>
<script lang="ts" setup>
import { ref, watchEffect } from "vue";
import { Atune } from "@/types/atune";
const shell = ref("here is every step's shell script");
let props = defineProps({
  selectedEditRow: {
    type: Object as () => Atune,
    default: null,
  },
});
let { prepare, restore, tune, workDir, tuneName, note, custom_name } =
  props.selectedEditRow;
const stepArr = ['开始', '环境准备', '调优', '环境恢复'];
const active = ref(1)
const next = () => {
  if (active.value++ > 3) active.value = 1
}
watchEffect(() => {
  switch (active.value) {
    case 2:
      shell.value = prepare;
      break;
    case 3:
      shell.value = tune;
      break;
    case 4:
      shell.value = restore;
      break;

    default:
      shell.value = "here is every step's shell script";
      break;
  }
})
</script>
<style lang="less" scoped>
.atuneDetail {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  background-color: #fff;

  &_con {
    width: 100%;
    height: 96%;
  }

  .shell {
    width: 90%;
    height: 100%;
    text-align: center;
  }
}
</style>
