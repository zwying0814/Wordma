<script setup lang="ts">
import {CommentDataType, CommentResponseType} from "@/types/Comment.ts";
import {alovaBaseUrlInstance} from "@/utils/http.ts";
import {message} from "ant-design-vue";
import {useUserStore} from "@/store/user.store.ts";

const openFlag = ref<boolean>(false);
const siteID = ref<string>();
const userStore = useUserStore();
const commentDataForm = ref<CommentDataType>({
  Content: "",
  UA: "",
  IP: "",
  Region: "",
  Type: "",
  Up: 0,
  Down: 0,
  PostTitle: "",
  PostID: 0,
  Username: "",
  UserID: 0,
  Parent: 0,
  ID: 0,
  CreatedAt: "",
});
const replyContent = ref<string>("");
const open = (comment: CommentResponseType | null, site: string) => {
  openFlag.value = true;
  if (comment)
    commentDataForm.value = {...comment, PostTitle: comment.Post.Slug, Username: comment.User.Name};
  if (siteID) siteID.value = site;
};

const close = () => {
  openFlag.value = false;
};

defineExpose({
  open,
  close,
});
const emits = defineEmits(["refresh"]);
const handleOk = async () => {
  await alovaBaseUrlInstance.Post(`/api/comment/`, {
    name: userStore.userInfo?.Name,
    content: replyContent.value,
    email: userStore.userInfo?.Email,
    url: userStore.userInfo?.Url,
    post_slug: commentDataForm.value.PostTitle,
    parent: commentDataForm.value.ID,
    site_id: parseInt(typeof siteID.value === "string" ? siteID.value : "0"),
  });
  message.success("回复成功");
  emits("refresh");
  close();
}
</script>

<template>
  <a-modal v-model:open="openFlag" title="快捷回复评论" @ok="handleOk">
    <a-form :model="commentDataForm" name="basic" layout="vertical" autocomplete="off">
      <a-alert class="my-4" :message="commentDataForm.Content" type="info"/>
      <a-form-item label="回复内容" name="Content" :rules="[{ required: true, message: '回复内容不能为空!' }]">
        <a-textarea v-model:value="replyContent"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>