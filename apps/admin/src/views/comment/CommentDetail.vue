<script setup lang="ts">
import {CommentDataType, CommentResponseType} from "@/types/Comment.ts";

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
</script>

<template>
  <a-drawer
      v-model:open="openFlag"
      :root-style="{ color: 'blue' }"
      style="color: red"
      title="评论详情"
      placement="right"
  >
    <a-form :model="commentDataForm" name="basic" layout="vertical" autocomplete="off">
      <a-form-item label="评论ID" name="ID" :rules="[{ required: true, message: '评论ID不能为空!' }]">
        <a-input disabled v-model:value="commentDataForm.ID" />
      </a-form-item>

      <a-form-item label="评论内容" name="Content" :rules="[{ required: true, message: '评论内容不能为空!' }]">
        <a-textarea disabled v-model:value="commentDataForm.Content" />
      </a-form-item>

      <a-form-item label="UA" name="UA">
        <a-input disabled v-model:value="commentDataForm.UA" />
      </a-form-item>

      <a-form-item label="IP" name="IP">
        <a-input disabled v-model:value="commentDataForm.IP" />
      </a-form-item>

      <a-form-item label="赞" name="Up">
        <a-input disabled v-model:value="commentDataForm.Up" />
      </a-form-item>

      <a-form-item label="踩" name="Down">
        <a-input disabled v-model:value="commentDataForm.Down" />
      </a-form-item>

      <a-form-item label="所属文章" name="PostTitle">
        <a-input disabled v-model:value="commentDataForm.PostTitle" />
      </a-form-item>

      <a-form-item label="评论者昵称" name="Username">
        <a-input disabled v-model:value="commentDataForm.Username" />
      </a-form-item>
    </a-form>
  </a-drawer>
</template>
