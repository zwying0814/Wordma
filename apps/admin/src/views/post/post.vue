<script setup lang="ts">
import Layout from "@/layout/layout.vue";
import {SelectProps} from "ant-design-vue";
import {useSitesStore} from "@/store/sites.store.ts";
import {alovaBaseUrlInstance} from "@/utils/http.ts";
import {PostListData, PostListDataResponse} from "@/types/Post.ts";

const columns = [
  {
    title: '文章Slug',
    dataIndex: 'slug',
    key: 'Slug',
  },
  {
    title: '阅读量',
    dataIndex: 'read',
    key: 'Read',
  },
  {
    title: '赞',
    dataIndex: 'up',
    key: 'Up',
  },
  {
    title: '踩',
    dataIndex: 'down',
    key: 'Down',
  }
];
const route = useRoute()
const sitesStore = useSitesStore()
onMounted(async () => {
  await sitesStore.loadSites()
  options.value = sitesStore.sites.map(i => {
    return {
      value: i.ID.toString(),
      label: i.Name
    }
  })
  if (route.query.siteID && options.value?.some(i => i.value == route.query.siteID)) {
    bindValue.value = route.query.siteID as string
  }
})
const bindValue = ref<string | undefined>(undefined);
const options = ref<SelectProps['options']>([]);
// 文章数据
const postsDataList = ref<PostListData[]>([])
const filterOption = (input: string, option: any) => {
  return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};

// 加载Post列表
const requestPosts = async (siteID: string) => {
  const res: PostListDataResponse = await alovaBaseUrlInstance.Get(`/api/post`, {
    params: {
      site_id: siteID,
      limit: 10,
      offset: 0
    }
  })
  console.log(res)
  postsDataList.value = res.data
}
watch(bindValue, () => {
  requestPosts(bindValue.value as string)
})
</script>

<template>
  <Layout>
    <template #header>
      <a-select
          v-model:value="bindValue"
          placeholder="选中一个站点"
          style="width: 200px"
          :options="options"
          :filter-option="filterOption"
      ></a-select>
    </template>
    <a-table :columns="columns" :data-source="postsDataList"></a-table>
  </Layout>
</template>