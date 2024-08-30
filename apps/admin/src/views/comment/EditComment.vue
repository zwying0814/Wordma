<script setup lang="ts">
import {CommentDataType, CommentResponseType} from "@/types/Comment.ts";
import {alovaBaseUrlInstance} from "@/utils/http.ts";
import {message} from "ant-design-vue";

const openFlag = ref<boolean>(false);
const commentDataForm = ref<CommentDataType>({
  Content: "",
  UA: "",
  IP: "",
  Region: "",
  Type: "",
  Up: 0,
  Down: 0,
  PostTitle:"",
  PostID: 0,
  Username: "",
  UserID: 0,
  Parent: 0,
  ID: 0,
  CreatedAt: "",
});
const open = (comment: CommentResponseType | null) => {
  openFlag.value = true;
  if (comment)
    commentDataForm.value = {...comment, PostTitle: comment.Post.Slug,Username: comment.User.Name};
};

const close = () => {
  openFlag.value = false;
};

defineExpose({
  open,
  close,
});
const emits = defineEmits(["refresh"]);
const handleOk =async () => {
  await alovaBaseUrlInstance.Put(`/api/comment/${commentDataForm.value.ID}`, commentDataForm.value);
  message.success("修改成功");
  emits("refresh");
  close();
}
</script>

<template>
  <a-modal v-model:open="openFlag" title="编辑内容" @ok="handleOk">
    <a-form :model="commentDataForm" name="basic" layout="vertical" autocomplete="off">
      <a-form-item label="评论内容" name="Content" :rules="[{ required: true, message: '评论内容不能为空!' }]">
        <a-textarea v-model:value="commentDataForm.Content" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>