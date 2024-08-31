<script setup>
import Sidebar from "@/layout/sidebar.vue";

const onCollapse = (collapsed, type) => {
  console.log(collapsed, type);
};
const onBreakpoint = broken => {
  console.log(broken);
};
</script>

<template>
  <a-layout class="w-screen h-screen">
    <a-layout-sider breakpoint="lg" class="!bg-white" collapsed-width="0" @collapse="onCollapse"
                    @breakpoint="onBreakpoint">
      <div class="logo"/>
      <Sidebar/>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="!px-4 !bg-white">
        <slot name="header"/>
      </a-layout-header>
      <a-layout-content :style="{ margin: '24px 16px 0' }">
        <slot/>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<style scoped>
#components-layout-demo-responsive .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}

.site-layout-sub-header-background {
  background: #fff;
}

.site-layout-background {
  background: #fff;
}

[data-theme='dark'] .site-layout-sub-header-background {
  background: #141414;
}
</style>