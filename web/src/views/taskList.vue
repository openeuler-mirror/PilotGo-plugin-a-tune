<template>
  <div class="taskList">
    <my-table ref="taskRef" :get-data="getTaskLists" :get-all-data="getTaskLists" :del-func="deleteTask"
      :search-func="searchTask">
      <template #listName>任务列表</template>
      <template #button_bar>
        <my-button @click="handleCreat">新增</my-button>
        <my-button @click="handleDelete">删除</my-button>
      </template>
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" label="编号" width="80" />
      <el-table-column prop="task_name" label="任务名称" />
      <el-table-column label="调优对象">
        <template #default="props">
          {{ props.row.tune.tuneName }}
        </template>
      </el-table-column>
      <el-table-column label="模板名称" :filters="allTune" :filter-method="filterTune">
        <template #default="props">
          <el-link type="primary" @click="atuneDetail(props.row)">{{ props.row.tune ? props.row.tune.custom_name : "暂无" }}
            <el-icon class="el-icon--right"><icon-view /></el-icon></el-link>
        </template>
      </el-table-column>
      <el-table-column prop="task_status" label="状态">
        <template #default="{ row }">
          <div style="display: flex;align-items: center; background-color: ;">
            <el-icon class="is-loading" color="#409eff" v-if="row.task_status == '运行中'">
              <Loading />
            </el-icon>
            <el-icon color="#67C23A" v-else-if="row.task_status == '完成'">
              <CircleCheck />
            </el-icon>
            <el-icon color="#E6A23C" v-else>
              <Clock />
            </el-icon>
            <span>{{ row.task_status }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="task_process" label="进度">
        <template #default="props">
          <el-progress :percentage="Math.floor(100 * (props.row.success_count / props.row.uuid_count))"><span>{{
            props.row.success_count
            + '/' + props.row.uuid_count }}</span></el-progress>
        </template>
      </el-table-column>
      <el-table-column prop="create_time" label="创建时间" />
      <el-table-column prop="update_time" label="更新时间" />
      <el-table-column label="操作" width="160">
        <template #default="{ row }">
          <my-button round size="small" :disabled="row.task_status === '运行中'" @click="handleStart(row)">{{ row.task_status
            === '等待' ? '启动' : '重启' }}</my-button>
          <my-button round size="small" @click="handleDetail(row)">详情</my-button>
        </template>
      </el-table-column>
    </my-table>
  </div>
  <el-dialog title="新增任务" width="50%" v-model="showDialog">
    <task-form @closeDialog="closeDialog" @refreshData="refreshData" />
  </el-dialog>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { View as IconView, Loading, CircleCheck, Clock } from "@element-plus/icons-vue";
import { getTaskLists, searchTask, deleteTask, getAtuneAllName, startTask } from "@/api/atune";
import { Task, ResultData } from "@/types/atune";
import { useRouter } from "vue-router";
import { useAtuneStore } from "@/store/atune";
import taskForm from "./taskForm.vue"
import { TableColumnCtx } from "element-plus";

const taskRef = ref();
const showDialog = ref(false);
interface FilterObj {
  text: string;
  value: string;
}
const allTune = ref<FilterObj[]>();
// 路由管理器
const router = useRouter();
const emit = defineEmits(["atuneDetail"]);

onMounted(() => {
  getAllTune();
})

// 获取所有的调优模板
const getAllTune = () => {
  getAtuneAllName().then((res) => {
    let tuneFilterArr: Array<FilterObj> = [];
    if (res.data.code === 200) {
      let tuneDatas = res.data.data;
      tuneDatas.map((tuneItem: string) => {
        tuneFilterArr.push({ text: tuneItem, value: tuneItem });
      });
      allTune.value = tuneFilterArr;

    }
  });
};

// 关闭dialog弹框
const closeDialog = () => {
  showDialog.value = false;
};

// 刷新表格数据
const refreshData = () => {
  taskRef.value.handleRefresh();
}

// 筛选模板任务
const filterTune = (value: string, row: Task, _column: TableColumnCtx<Task>) => {
  console.log(value, row)
  return row.tune.custom_name === value;
}

// 开始执行任务
const handleStart = (row: Task) => {
  let params = {
    machine_uuids: row.machine_uuids,
    tuneId: row.tune_id,
    taskId: row.id
  }
  startTask({ ...params }).then((res: { data: ResultData }) => {
    if (res.data.code === 200) {
      taskRef.value.handleRefresh();
    }
  })
}

// 查看模板详情
const atuneDetail = (row: Task) => {
  emit("atuneDetail", row);
};
// 查看任务详情
const handleDetail = (row: Task) => {
  router.push("/task/detail");
  useAtuneStore().setTaskRow(row);
};


// 新增
const handleCreat = () => {
  showDialog.value = true;
};
// 删除
const handleDelete = () => {
  taskRef.value.handleDelete();
};
</script>

<style lang="less" scoped>
.taskList {
  width: 100%;
  height: 100%;
}
</style>
