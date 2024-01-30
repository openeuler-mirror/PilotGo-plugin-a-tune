<template>
  <div>
    <el-form :model="form" class="custom-form">
      <el-form-item label="任务名称">
        <el-input v-model="form.task_name"></el-input>
      </el-form-item>
      <el-form-item label="下发机器">
        <el-select v-model="form.macs" placeholder="请选择下发机器" multiple>
          <el-option v-for="item in allMac" :key="item" :label="item.ip" :value="item.uuid" />
        </el-select>
      </el-form-item>
      <el-form-item label="调优对象">
        <el-select v-model="form.tune_id" placeholder="请选择调优模板">
          <el-option v-for="item in allTune" :key="item" :label="item.custom_name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="任务描述">
        <el-input v-model="form.description" type="textarea"></el-input>
      </el-form-item>
    </el-form>
    <el-form class="centered-buttons">
      <my-button @click="onSubmit">保存</my-button>
      <my-button @click="onCancel">取消</my-button>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ElMessage } from "element-plus";
import { reactive, onMounted, ref } from "vue";
import { addTask, getAtuneAllList, getAllMachine } from "@/api/atune";
import { Atune } from "@/types/atune";
interface MacItem {
  uuid: string;
  ip: string;
  [propName: string]: unknown;
}
const emit = defineEmits(["closeDialog", 'refreshData']);
const allTune = ref([] as Atune[]);
let allMac = reactive([] as MacItem[]);
let uuidArr = ref(['']);
const form = reactive({
  task_name: "",
  macs: "",
  tune_id: "",
  description: ""
});

onMounted(() => {
  getAllTune();
  getMacs();
})

// 获取调优模板列表
const getAllTune = () => {
  getAtuneAllList().then((res) => {
    if (res.data.code === 200) {
      allTune.value = res.data.data;
    }
  });
};
// 获取所有机器列表
const getMacs = () => {
  getAllMachine().then((res) => {
    if (res.data.code === 200) {
      allMac = res.data.data;
    }
  })
}


// 保存任务数据
const saveTaskData = () => {
  let params = {
    machine_uuids: form.macs,
    tune_id: form.tune_id,
    task_name: form.task_name,
    description: form.description
  }
  addTask({ ...params })
    .then((res) => {
      if (res.data.code === 200) {
        ElMessage.success(res.data.msg);
        emit('refreshData');
      } else {
        ElMessage.error(res.data.msg);
      }
    })
    .catch((err) => {
      ElMessage.error("数据传输失败，请检查", err);
    });
};

const onSubmit = () => {
  emit("closeDialog");
  saveTaskData();
};

const onCancel = () => {
  emit("closeDialog");
};
</script>

<style scoped lang="less">
.custom-form {
  margin-left: 25px;
  margin-right: 20px;
  outline-width: 120px;

  .custom-input {
    white-space: pre-wrap;
    resize: none;
    text-align: left;
    vertical-align: top;
  }
}

.centered-buttons {
  text-align: center;
}
</style>