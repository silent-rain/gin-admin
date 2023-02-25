<template>
  <el-form
    v-if="configHash.website_title"
    :model="configHash"
    label-width="100px"
    style="width: 100%"
  >
    <el-form-item label="角色名称" prop="website_title">
      <el-input
        v-model="configHash.website_title.value"
        placeholder="请输入角色名称"
      />
    </el-form-item>
    <el-form-item label="备注" prop="website_intro">
      <el-input
        v-model="configHash.website_intro.value"
        placeholder="请输入备注"
      />
    </el-form-item>
  </el-form>
  <div class="submit">
    <el-button type="primary" @click="submitForm">保存变更</el-button>
  </div>
</template>

<script setup lang="ts">
import { ElMessage } from 'element-plus';
import {
  getConfigChildrenByKey,
  batchUpdateConfig,
} from '@/api/data-center/config';
import { ConfigListRsp, WebsiteConfig } from '@/typings/api/data-center/config';
import { WebsiteSettings } from '@/constant/data-center/config';

const configHash = ref<WebsiteConfig>({} as WebsiteConfig);

onBeforeMount(() => {
  fetchConfigChildrenByKey();
});

// 通过上级 key 获取子配置列表
const fetchConfigChildrenByKey = async () => {
  try {
    const resp = (
      await getConfigChildrenByKey({
        key: WebsiteSettings,
      })
    ).data as ConfigListRsp;
    for (const item of resp.data_list) {
      configHash.value[item.key] = item;
      console.log(item.key);
    }
  } catch (error) {
    console.log(error);
  }
};

// 提交
const submitForm = async () => {
  try {
    await batchUpdateConfig({});
    ElMessage.success('操作成功');
  } catch (error) {
    console.log(error);
  }
};
</script>

<style scoped lang="scss">
.submit {
  text-align: center;
}
</style>
