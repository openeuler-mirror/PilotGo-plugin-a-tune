<template>
  <div class="task_detail">
    <el-icon class="task_detail_back" size="20" @click="goBack">
      <Back />
    </el-icon>
    <el-descriptions :column="1">
      <el-descriptions-item label="任务名称：">
        {{ taskItem.task_name }}</el-descriptions-item>
      <el-descriptions-item label="创建时间：">
        {{ taskItem.create_time }}</el-descriptions-item>
      <el-descriptions-item label="任务进度：" style="display:flex; align-items: center;">
        <el-progress style="width: 30%;"
          :percentage="Math.floor(100 * (taskItem.success_count / taskItem.uuid_count))"><span>{{
      taskItem.success_count
      + '/' + taskItem.uuid_count }}</span></el-progress>
      </el-descriptions-item>
      <el-descriptions-item label="任务描述：">{{
      taskItem.description
    }}</el-descriptions-item>
    </el-descriptions>
    <el-divider content-position="left">调优机器进度&nbsp;
      <el-button size="small" plain type="primary" @click="getDetail">刷新</el-button>
    </el-divider>
    <div class="mac" v-if="showMacProcess">
      <div class="mac_li" v-for="(item, index) in uuids" :key="item.machineUUID"
        @click="handleCompare(item.machineUUID)">
        <span class="mac_li_uuid">
          <p><el-icon>
              <Monitor />
            </el-icon>
            <span>:&nbsp;{{ item.machineIp }}</span>
          </p>
          <p>cpu架构：{{ item.cpuArch }}</p>
          <p>操作系统：{{ item.os }}</p>
        </span>
        <el-timeline class="time">
          <el-timeline-item :hollow="true" type="primary" size="large">
            开始
          </el-timeline-item>
          <el-timeline-item v-for="(stepItem, stepIndex) in item.result" :key="stepIndex"
            :icon="handleStepTimeLine(stepItem.is_success).icon" :color="handleStepTimeLine(stepItem.is_success).color"
            size="large" :timestamp="stepItem.startTime">
            <div>
              <p class="time_step" @click.stop="handleCommand(stepItem.stdout, index)">{{ stepItem.command_type }}</p>
            </div>
          </el-timeline-item>
        </el-timeline>
      </div>
    </div>
  </div>
  <el-drawer v-model="drawer" title="性能提升看板" direction="ttb" size="80%">
    <tuneResult :tune_uuid="currentUuid" :taskId="taskItem.id" />
  </el-drawer>
  <el-drawer v-model="stepDrawer" title="步骤输出详情" direction="rtl" size="60%">
    <el-scrollbar class="scroll">
      <pre>{{ stepCommand }}</pre>
    </el-scrollbar>
  </el-drawer>
</template>
<script lang="ts" setup>
import { onMounted, reactive, ref, toRefs, watchEffect } from "vue";
import { getTaskDetail } from "@/api/result";
import { MoreFilled, Monitor, Check, Close, Loading, Back } from '@element-plus/icons-vue'
import { ResultData, } from "@/types/atune";
import { useAtuneStore } from "@/store/atune";
import tuneResult from "./task/tuneResult.vue";
import router from "@/router";

interface MacItem {
  taskId: number;
  machineUUID: string;
  machineIp: string;
  [propName: string]: any;
}
let taskItem: any = reactive({
  description: '',
  task_name: '',
  create_time: '',
  success_count: null,
  id: null,
  uuid_count: null

});
let uuids = reactive([] as MacItem[]);
const showMacProcess = ref(false);
const drawer = ref(false);
const stepDrawer = ref(false);
const stepCommand = ref('');
const currentUuid = ref('');

onMounted(() => {
  getDetail();
})
const goBack = () => {
  router.push('/')
}
const getDetail = () => {
  showMacProcess.value = false;
  uuids = [];
  getTaskDetail({ taskId: taskItem.id }).then((res: { data: ResultData }) => {
    if (res.data.code === 200) {
      uuids = res.data.data;
      showMacProcess.value = true;
    }
  })
}

// 处理调优结果对比
const handleCompare = (uuid: string) => {
  drawer.value = true;
  currentUuid.value = uuid;
}

// 处理输出
const handleCommand = (outString: string, _index: number) => {
  stepDrawer.value = true;
  stepCommand.value = outString;
}

// 处理时间线的图标和颜色
const handleStepTimeLine = (is_success: string) => {
  let stepObj: {
    icon: any;
    color: string;
  } = { icon: '', color: '' }
  switch (is_success) {
    case 'waiting':
      stepObj.color = '#E6A23C';
      stepObj.icon = MoreFilled;
      break;
    case 'running':
      stepObj.color = '#409EFF';
      stepObj.icon = Loading;
      break;
    case 'fail':
      stepObj.color = '#F56C6C';
      stepObj.icon = Close;
      break;

    default:
      stepObj.color = '#67C23A';
      stepObj.icon = Check;
      break;
  }
  return stepObj;
}

watchEffect(() => {
  let { description, task_name, create_time, id, success_count, uuid_count } = toRefs(useAtuneStore().taskRow);
  taskItem.description = description;
  taskItem.create_time = create_time;
  taskItem.task_name = task_name;
  taskItem.id = id;
  taskItem.success_count = success_count;
  taskItem.uuid_count = uuid_count;
})

</script>
<style lang="less" scope>
.task_detail {
  width: 98%;
  height: calc(100% - 20px);
  padding: 20px;
  margin: 0 auto;
  background-color: #fff;
  border-radius: 6px;

  &_back {
    cursor: pointer;

    &:hover {
      color: var(--active-color)
    }
  }


  .mac {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    height: calc(100% - 60px - 20px - 120px);
    overflow-y: scroll;

    &_li {
      width: 380px;
      height: 400px;
      margin: 10px 16px;
      display: flex;
      justify-content: center;
      border-radius: 4px;
      align-items: center;
      flex-wrap: wrap;
      border: 1px solid #dcdfe6;
      transition: all .5s;
      position: relative;
      overflow: hidden;

      &_uuid {
        display: flex;
        width: 100%;
        align-items: center;
        justify-content: center;
        flex-wrap: wrap;
        color: var(--el-color-info-dark-2);

        p {
          display: flex;
          align-items: center;
          justify-content: flex-start;
          width: 80%;
        }
      }

      .time {
        width: 46%;

        &_step {
          &:hover {
            color: #409EFF;
          }
        }
      }

      &:hover {
        cursor: pointer;
        box-shadow: var(--el-box-shadow-dark);
      }


    }
  }
}
</style>
