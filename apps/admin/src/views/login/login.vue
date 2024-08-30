<script setup lang="ts">
import {message} from "ant-design-vue";
import {useUserStore} from "@/store/user.store.ts";

const loading = ref(false);
const router = useRouter();
const userStore = useUserStore();

const loginFormModel = ref({
  username: 'admin',
  password: '123456',
});

const handleSubmit = async () => {
  const { username, password } = loginFormModel.value;
  if (username.trim() == '' || password.trim() == '') {
    return message.warning('用户名或密码不能为空！');
  }
  loading.value = true;
  await userStore.login(loginFormModel.value);
  message.success('登录成功！');
  setTimeout(() => router.push('/'));
  loading.value = false;
};
</script>

<template>
  <div class="login-box flex flex-col items-center w-screen h-screen justify-center">
      <h1 class="text-3xl font-bold mb-6">📝 Wordma</h1>
    <a-form layout="horizontal" :model="loginFormModel" @submit.prevent="handleSubmit" class="w-72">
      <a-form-item>
        <a-input v-model:value="loginFormModel.username" size="large" placeholder="admin">
          <template #prefix> <UserOutlined /> </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-input
            v-model:value="loginFormModel.password"
            size="large"
            type="password"
            placeholder="a123456"
            autocomplete="new-password"
        >
          <template #prefix> <LockOutlined /></template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-button type="primary" html-type="submit" size="large" :loading="loading" block>
          登录
        </a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
