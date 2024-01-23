import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import MyTable from "@/components/table.vue";
import MyButton from "@/components/myButton.vue";
import { ElDialog } from "element-plus";
import "element-plus/dist/index.css";
import router from "./router";
import pinia from '@/store'

const app = createApp(App);
// 设置全局变量
// app.config.globalProperties.echarts = echarts;
app.component("el-dialog", ElDialog);
app.component("my-table", MyTable);
app.component("my-button", MyButton);

app.use(pinia);
app.use(router);
app.use(ElementPlus);

app.mount("#app");
