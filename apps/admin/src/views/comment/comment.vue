<script setup lang="ts">
import Layout from "@/layout/layout.vue";
import {message, SelectProps} from 'ant-design-vue';
import {useSitesStore} from "@/store/sites.store.ts";
import {alovaBaseUrlInstance} from "@/utils/http.ts";
import {CommentResponse, CommentDataType, CommentResponseType} from "@/types/Comment.ts";
import CommentDetail from "@/views/comment/CommentDetail.vue";
import {AntTableCell} from "@/types/ant.ts";
import EditComment from "@/views/comment/EditComment.vue";
import ReplyComment from "@/views/comment/ReplyComment.vue";

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
const options = ref<SelectProps['options']>([]);
const bindValue = ref<string | undefined>(undefined);
const filterOption = (input: string, option: any) => {
  return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};

// 评论数据
const commentsDataList = ref<CommentDataType[]>([])
// 加载评论列表
const requestComments = async (siteID: string) => {
  const res: CommentResponse = await alovaBaseUrlInstance.Get(`/site/comments`, {
    params: {
      site_id: siteID,
      limit: 10,
      offset: 0,
      sort_by: 'desc',
    }
  })
  commentsDataList.value = res.data
  console.log(commentsDataList.value)
}
watch(bindValue, () => {
  requestComments(bindValue.value as string)
})

const columns = [
  {
    title: '评论ID',
    dataIndex: 'ID',
    key: 'ID',
  },
  {
    title: '评论内容',
    dataIndex: 'Content',
    key: 'Content',
  },
  {
    title: '所属文章Slug',
    dataIndex: ['Post', 'Slug'],
    key: 'Post.Slug',
  },
  {
    title: '时间',
    key: 'CreatedAt',
    dataIndex: 'CreatedAt',
  },
  {
    title: '评论作者',
    dataIndex: ['User', 'Name'],
    key: 'User.Name',
  },
  {
    title: '状态',
    key: 'Type',
    dataIndex: 'Type',
  },
  {
    title: '操作',
    key: 'Options',
    dataIndex: 'Options',
  },

];

const typeConvert = (type: string) => {
  switch (type) {
    case 'published':
      return "公开"
    case 'pending':
      return '待审核'
    case 'trash':
      return '已删除'
    default:
      return '未知'
  }
}

// 通过评论
const passComment = async (commentID: number) => {
  await alovaBaseUrlInstance.Put(`/comment/${commentID}`, {
    type: 'published'
  })
  message.success("评论已通过")
  await refresh()
}
// 删除评论
const deleteComment = async (commentID: number) => {
  await alovaBaseUrlInstance.Delete(`/comment/${commentID}`)
  message.success("评论已删除")
  await refresh()
}
const refresh = async () => {
  await requestComments(bindValue.value as string)
}


const commentDetailRef = ref<InstanceType<typeof CommentDetail> | null>(null)
const viewCommentDetail = (comment: CommentResponseType) => {
  if (commentDetailRef.value) commentDetailRef.value.open(comment)
}

const editCommentRef = ref<InstanceType<typeof EditComment> | null>(null)
const editCommentContent = (comment: CommentResponseType) => {
  if (editCommentRef.value) editCommentRef.value.open(comment)
}

const replyCommentRef = ref<InstanceType<typeof ReplyComment> | null>(null)
const replyComment = (comment: CommentResponseType) => {
  if (replyCommentRef.value) replyCommentRef.value.open(comment, bindValue.value as string)
}
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
    <a-table :columns="columns" :data-source="commentsDataList">

      <template #bodyCell="{ column, record }:AntTableCell">

        <template v-if="column.key === 'Type'">
          <a-tag color="success" v-if="record.Type==='published'">
            {{ typeConvert(record.Type) }}
          </a-tag>
          <a-tag color="warning" v-else-if="record.Type==='pending'">
            {{ typeConvert(record.Type) }}
          </a-tag>
          <a-tag color="error" v-else-if="record.Type==='trash'">
            {{ typeConvert(record.Type) }}
          </a-tag>
          <a-tag v-else>
            {{ typeConvert(record.Type) }}
          </a-tag>
        </template>

        <template v-if="column.key==='Options'">
          <a-flex :gap="10">
            <a-button type="primary" size="small" @click="passComment(record.ID)">通过</a-button>
            <a-button size="small" @click="viewCommentDetail(record)">详情</a-button>
            <a-button type="primary" size="small" @click="replyComment(record)">回复</a-button>
            <a-button size="small" @click="editCommentContent(record)">编辑</a-button>
            <a-popconfirm placement="topRight"
                          title="确定删除评论吗？该评论下的子级评论将会清空！"
                          ok-text="确认删除"
                          cancel-text="取消"
                          @confirm="deleteComment(record.ID)"
            >
              <a-button danger size="small">删除</a-button>
            </a-popconfirm>

          </a-flex>
        </template>
      </template>
    </a-table>
  </Layout>
  <CommentDetail ref="commentDetailRef"/>
  <EditComment ref="editCommentRef" @refresh="refresh"/>
  <ReplyComment ref="replyCommentRef" @refresh="refresh"/>
</template>