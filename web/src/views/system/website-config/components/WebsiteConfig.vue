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
    <el-form-item
      :label="configHash.website_title_brief.name"
      prop="website_title_brief"
    >
      <el-input
        v-model="configHash.website_title_brief.value"
        :placeholder="`请输入${configHash.website_title_brief.name}`"
      />
    </el-form-item>

    <el-form-item :label="configHash.website_logo.name" prop="website_logo">
      <UploadLogo v-model:url="configHash.website_logo.value"></UploadLogo>
    </el-form-item>
    <el-form-item
      :label="configHash.website_propaganda.name"
      prop="website_propaganda"
    >
      <PhotoWall v-model:file-list="websitePropagandas"></PhotoWall>
    </el-form-item>

    <el-divider />

    <!-- 网站SEO设置 -->
    <el-form-item
      :label="configHash.website_description.name"
      prop="website_description"
    >
      <el-input
        v-model="configHash.website_description.value"
        type="textarea"
        :placeholder="`请输入${configHash.website_description.name}`"
      />
    </el-form-item>
    <el-form-item
      :label="configHash.website_keywords.name"
      prop="website_keywords"
    >
      <el-select
        v-model="state.websiteKeywords"
        multiple
        filterable
        allow-create
        no-match-text
        :reserve-keyword="false"
        :placeholder="`请输入${configHash.website_keywords.name}`"
      >
        <el-option
          v-for="item in websiteKeywordsOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
    <el-form-item :label="configHash.website_tags.name" prop="website_tags">
      <el-select
        v-model="state.websiteTags"
        multiple
        filterable
        allow-create
        no-match-text
        :reserve-keyword="false"
        :placeholder="`请输入${configHash.website_tags.name}`"
      >
        <el-option
          v-for="item in websiteTagsOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-form-item>
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
        type="textarea"
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
import UploadLogo from '@/components/Upload/UploadLogo.vue';
import PhotoWall from '@/components/Upload/PhotoWall.vue';

const configHash = ref<WebsiteConfig>({} as WebsiteConfig);
const state = reactive({
  websiteTags: [] as string[],
  websiteKeywords: [] as string[],
});
const websiteTagsOptions = computed(() => {
  if (!configHash.value.website_tags.value) {
    return [];
  }
  const tags = configHash.value.website_tags.value.split(',');
  state.websiteTags = [...tags];
  return tags.map((v) => {
    return {
      label: v,
      value: v,
    };
  });
});
const websiteKeywordsOptions = computed(() => {
  if (!configHash.value.website_keywords.value) {
    return [];
  }
  const keywords = configHash.value.website_keywords.value.split(',');
  state.websiteKeywords = [...keywords];
  return keywords.map((v) => {
    return {
      label: v,
      value: v,
    };
  });
});

const websitePropagandas = computed({
  get() {
    if (!configHash.value.website_propaganda.value) {
      return [];
    }
    const list = JSON.parse(configHash.value.website_propaganda.value);
    return list;
  },
  set(val) {
    configHash.value.website_propaganda.value = JSON.stringify(val);
  },
});

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
    }
  } catch (error) {
    console.log(error);
  }
};

// 提交
const submitForm = async () => {
  const data = [] as any[];
  for (const k in configHash.value) {
    let value = configHash.value[k].value;
    if (k == 'website_tags') {
      value = state.websiteTags.join(',');
    } else if (k == 'website_keywords') {
      value = state.websiteKeywords.join(',');
    }
    data.push({
      id: configHash.value[k].id,
      name: configHash.value[k].name,
      key: k,
      value: value,
    });
  }

  try {
    await batchUpdateConfig(data);
    await fetchConfigChildrenByKey();
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
