<script setup lang="ts">

import {MenuClickEventHandler} from "ant-design-vue/es/menu/src/interface";

const route = useRoute()
const selectedKeys = ref([route.path]);
const router = useRouter();
const menuClick:MenuClickEventHandler = (info) => {
  router.push(info.key as string)
}
</script>

<template>
  <a-menu v-model:selectedKeys="selectedKeys" theme="light" mode="inline" @click="menuClick">
    <a-menu-item key="/home">
      <user-outlined/>
      <span class="nav-text">站点列表</span>
    </a-menu-item>
    <a-menu-item key="/comment">
      <user-outlined/>
      <span class="nav-text">评论管理</span>
    </a-menu-item>
    <a-menu-item key="/post">
      <user-outlined/>
      <span class="nav-text">文章数据</span>
    </a-menu-item>
  </a-menu>
</template>