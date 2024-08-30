<script setup lang="ts">
import Layout from "@/layout/layout.vue";
import {useSitesStore} from "@/store/sites.store.ts";
import AddSite from "@/views/home/SiteDrawer.vue";
import {message} from "ant-design-vue";
import {Site} from "@/types/Site.ts";

const siteDrawer = ref<InstanceType<typeof AddSite> | null>(null)
const sitesStore = useSitesStore()
const router = useRouter()
onMounted(() => {
  sitesStore.loadSites()
})

const openAddSiteDrawer = () => {
  if (siteDrawer.value)
    siteDrawer.value.open(null)
}

const deleteSite = async (id: number) => {
  await sitesStore.deleteSite(id)
  message.success('删除成功')
}
const editSite = (site:Site) => {
  if (siteDrawer.value)
    siteDrawer.value.open(site)
}
</script>

<template>
  <Layout>
    <template #header>
      <a-button type="primary" @click="openAddSiteDrawer">新增站点</a-button>
    </template>
    <a-row :gutter="16">
      <a-col :xl="6" :lg="8" :sm="12" :xs="24" v-for="item in sitesStore.sites" :key="item.ID">
        <a-card :title="item.Name" :bordered="false">
          <p>地址：{{ item.Url }}</p>
          <template #extra>
            <a-button class="mr-2" size="small" @click="editSite(item)">编辑</a-button>
            <a-popconfirm title="确定删除站点吗？全部数据将会清空！" ok-text="确认删除" cancel-text="取消" @confirm="deleteSite(item.ID)">
              <a-button type="primary" danger size="small">删除</a-button>
            </a-popconfirm>
          </template>
          <template #actions>
            <span @click="router.push(`/comment?siteID=${item.ID}`)">评论</span>
            <span @click="router.push(`/post?siteID=${item.ID}`)">文章数据</span>
          </template>
        </a-card>
      </a-col>
    </a-row>
  </Layout>
  <AddSite ref="siteDrawer" @refresh="sitesStore.loadSites()"/>
</template>