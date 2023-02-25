<template>
  <el-form
    v-if="configHash.website_title"
    :model="configHash"
    label-width="100px"
    style="width: 100%"
  >
    <!-- 网站设置 -->
    <el-form-item :label="configHash.website_title.name" prop="website_title">
      <el-input
        v-model="configHash.website_title.value"
        :placeholder="`请输入${configHash.website_title.name}`"
      />
    </el-form-item>
    <el-form-item :label="configHash.website_intro.name" prop="website_intro">
      <el-input
        v-model="configHash.website_intro.value"
        :placeholder="`请输入${configHash.website_intro.name}`"
      />
    </el-form-item>
    <el-form-item
      :label="configHash.website_keyword.name"
      prop="website_keyword"
    >
      <el-input
        v-model="configHash.website_keyword.value"
        :placeholder="`请输入${configHash.website_keyword.name}`"
      />
    </el-form-item>
    <el-form-item :label="configHash.website_desc.name" prop="website_desc">
      <el-input
        v-model="configHash.website_desc.value"
        :placeholder="`请输入${configHash.website_desc.name}`"
      />
    </el-form-item>
    <el-form-item :label="configHash.website_tags.name" prop="website_tags">
      <el-input
        v-model="configHash.website_tags.value"
        :placeholder="`请输入${configHash.website_tags.name}`"
      />
    </el-form-item>

    <el-form-item :label="configHash.website_logo.name" prop="website_logo">
      <el-input
        v-model="configHash.website_logo.value"
        placeholder="请输入网站LOGO"
      />
    </el-form-item>
    <el-form-item
      :label="configHash.website_propaganda.name"
      prop="website_propaganda"
    >
      <el-input
        v-model="configHash.website_propaganda.value"
        placeholder="请输入网站宣传片"
      />
    </el-form-item>

    <el-divider />

    <!-- 网站SEO设置 -->
    <el-form-item
      :label="configHash.website_seo_title.name"
      prop="website_title"
    >
      <el-input
        v-model="configHash.website_seo_title.value"
        :placeholder="`请输入${configHash.website_seo_title.name}`"
      />
    </el-form-item>
    <el-form-item
      :label="configHash.website_seo_desc.name"
      prop="website_title"
    >
      <el-input
        v-model="configHash.website_seo_desc.value"
        :placeholder="`请输入${configHash.website_seo_desc.name}`"
      />
    </el-form-item>

    <el-divider />

    <!-- 网站联系方式 -->
    <el-form-item
      :label="configHash.website_company_address.name"
      prop="website_company_address"
    >
      <el-input
        v-model="configHash.website_company_address.value"
        :placeholder="`请输入${configHash.website_company_address.name}`"
      />
    </el-form-item>
    <el-form-item :label="configHash.website_phone.name" prop="website_phone">
      <el-input
        v-model="configHash.website_phone.value"
        :placeholder="`请输入${configHash.website_phone.name}`"
      />
    </el-form-item>
    <el-form-item :label="configHash.website_email.name" prop="website_email">
      <el-input
        v-model="configHash.website_email.value"
        :placeholder="`请输入${configHash.website_email.name}`"
      />
    </el-form-item>
    <el-form-item :label="configHash.website_qq.name" prop="website_qq">
      <el-input
        v-model="configHash.website_qq.value"
        :placeholder="`请输入${configHash.website_qq.name}`"
      />
    </el-form-item>

    <!-- 版权/备案号 -->
    <el-form-item
      :label="configHash.website_copyright.name"
      prop="website_copyright"
    >
      <el-input
        v-model="configHash.website_copyright.value"
        :placeholder="`请输入${configHash.website_copyright.name}`"
      />
    </el-form-item>
    <el-form-item
      :label="configHash.website_filing_number.name"
      prop="website_filing_number"
    >
      <el-input
        v-model="configHash.website_filing_number.value"
        :placeholder="`请输入${configHash.website_filing_number.name}`"
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
