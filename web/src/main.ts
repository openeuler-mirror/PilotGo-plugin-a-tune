/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
 */
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
app.component("el-dialog", ElDialog);
app.component("my-table", MyTable);
app.component("my-button", MyButton);

app.use(pinia);
app.use(router);
app.use(ElementPlus);

app.mount("#app");
