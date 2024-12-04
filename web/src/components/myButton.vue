<!--
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugin-a-tune licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jan 12 14:12:37 2024 +0800
-->
<template>
  <el-button class="button" v-bind="$attrs" plain>
    <slot></slot>
  </el-button>
</template>
<script lang="ts" setup></script>
<style lang="less" scoped>
.button {
  color: rgb(96, 122, 207) !important;
  &:hover {
    background-color: #fff !important;
    color: rgb(96, 122, 207) !important;
    border: 1px solid rgb(96, 122, 207) !important;
  }
}
</style>
