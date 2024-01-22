<template>
  <div class="tuneList shadow">
    <my-table ref="tuneRef" :get-data="getTuneLists" :get-all-data="getTuneLists" :del-func="deleteTune"
      :search-func="searchTune">
      <template #listName>模板列表</template>
      <template #button_bar>
        <my-button @click="handleCreat">新增</my-button>
        <my-button @click="handleDelete">删除</my-button>
      </template>
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="编号" width="100" />
      <el-table-column prop="tuneName" label="调优对象" width="100" />
      <el-table-column prop="custom_name" label="自定义名称" width="100" />
      <el-table-column prop="description" label="概述" />
      <el-table-column prop="create_time" label="创建时间" />
      <el-table-column prop="update_time" label="更新时间" />
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <my-button size="small" @click="handleDetail(row)">详情</my-button>
          <my-button size="small" @click="handleEdit(row)">编辑</my-button>
        </template>
      </el-table-column>
    </my-table>
  </div>
  <el-dialog :title="title" :width="width" v-model="showDialog">
    <atuneTemplete v-if="type === 'edit'" :is-tune="true" :is-update="isUpdate" :selectedEditRow="selectedEditRow"
      @closeDialog="closeDialog" @refresh-data="handleRefresh">
    </atuneTemplete>
    <tuneDetail v-if="type === 'detail'" :selectedEditRow="selectedEditRow" @closeDialog="closeDialog" />
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { getTuneLists, searchTune, deleteTune } from "@/api/atune";
import atuneTemplete from "@/components/atuneTemplete.vue";
import { Atune } from "@/types/atune";
import tuneDetail from "./atuneDetail.vue"
const tuneRef = ref();
const showDialog = ref(false);
const selectedEditRow = ref({} as Atune);
const type = ref(''); // 弹窗类型
const title = ref('调优模板信息');
const width = ref('50%')
const isUpdate = ref(false)

// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
};
// 新增
const handleCreat = () => {
  showDialog.value = true;
  type.value = 'edit';
  isUpdate.value = false;
};
// 删除
const handleDelete = () => {
  tuneRef.value.handleDelete();
};
// 刷新
const handleRefresh = () => {
  tuneRef.value.handleRefresh();
}
// 详情
const handleDetail = (row: Atune) => {
  showDialog.value = true;
  selectedEditRow.value = row;
  type.value = 'detail';
  title.value = row.custom_name + '详情';
  width.value = '70%';
};
// 编辑
const handleEdit = (row: Atune) => {
  selectedEditRow.value = row;
  showDialog.value = true;
  type.value = 'edit';
  isUpdate.value = true;
  title.value = row.custom_name + '信息';
};
</script>

<style lang="less" scoped>
.tuneList {
  width: 98%;
  margin: 0 auto;
  height: calc(100% - 44px - 10px);
}
</style>
