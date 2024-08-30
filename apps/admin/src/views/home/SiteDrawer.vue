<script lang="ts" setup>
import {alovaBaseUrlInstance} from "@/utils/http.ts";
import {message} from "ant-design-vue";
import {Site} from "@/types/Site.ts";

const openFlag = ref<boolean>(false);
// 编辑or新增
const isEdit = ref<boolean>(false);
// 编辑的站点
const open = (site: Site | null) => {
  openFlag.value = true;
  isEdit.value = !!site;
  if (isEdit.value && site) {
    addSiteForm.value = {
      id: site.ID,
      name: site.Name,
      url: site.Url,
    }
  }else{
    addSiteForm.value = {
      id: 0,
      name: '',
      url: '',
    }
  }
}
const close = () => {
  openFlag.value = false;
}
defineExpose({
  open,
  close
})

const addSiteForm = ref({
  id: 0,
  name: '',
  url: '',
});

const loading = ref<boolean>(false);

const onFinish = async (values: any) => {
  loading.value = true;
  if (isEdit.value) {
    await alovaBaseUrlInstance.Put(`/api/site/${values.id}`, values)
    message.success('编辑成功');
  } else {
    await alovaBaseUrlInstance.Post('/api/site', values)
    message.success('添加成功');
  }
  close();
  emits('refresh');
  loading.value = false;
};

const onFinishFailed = (errorInfo: any) => {
  console.log('Failed:', errorInfo);
};

const emits = defineEmits(['refresh']);
</script>

<template>
  <a-drawer v-model:open="openFlag" :title="`${isEdit?'编辑':'新增'}站点`" placement="right">
    <a-form
        :model="addSiteForm"
        name="basic" layout="vertical"
        autocomplete="off"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
    >
      <a-form-item v-if="isEdit"
          label="站点ID" name="id" :rules="[{ required: true, message: '站点ID不能为空!' }]">
        <a-input disabled v-model:value="addSiteForm.id"/>
      </a-form-item>
      <a-form-item
          label="站点名称" name="name" :rules="[{ required: true, message: '请输入站点名称!' }]">
        <a-input v-model:value="addSiteForm.name"/>
      </a-form-item>

      <a-form-item label="站点Url" name="url" :rules="[{ required: true, message: '请输入站点Url，无需携带协议头' }]">
        <a-input v-model:value="addSiteForm.url"/>
      </a-form-item>


      <a-button class="w-full mt-4" block type="primary" html-type="submit" :loading="loading">提交</a-button>

    </a-form>
  </a-drawer>
</template>


