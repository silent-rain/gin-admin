<template>
  <div class="footer">
    <el-form
      v-if="configHash.website_title"
      :model="configHash"
      label-width="100px"
      style="width: 100%"
      size="small"
    >
      <!-- 网站联系方式 -->
      <el-row justify="center">
        <el-col :span="6">
          <el-form-item
            :label="configHash.website_company_address.name"
            prop="website_company_address"
          >
            {{ configHash.website_company_address.value }}
          </el-form-item>
          <el-form-item
            :label="configHash.website_phone.name"
            prop="website_phone"
          >
            {{ configHash.website_phone.value }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item
            :label="configHash.website_email.name"
            prop="website_email"
          >
            {{ configHash.website_email.value }}
          </el-form-item>
          <el-form-item :label="configHash.website_qq.name" prop="website_qq">
            {{ configHash.website_qq.value }}
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <!-- 版权/备案号 -->
          <el-form-item
            :label="configHash.website_copyright.name"
            prop="website_copyright"
          >
            {{ configHash.website_copyright.value }}
          </el-form-item>
          <el-form-item
            :label="configHash.website_filing_number.name"
            prop="website_filing_number"
          >
            {{ configHash.website_filing_number.value }}
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { useBasicStore } from '@/store/basic';
import { WebsiteConfig } from '@/typings/api/data-center/config';

const basicStore = useBasicStore();

const configHash = computed(() => {
  console.log(basicStore.webSiteConfigMap);
  return basicStore.webSiteConfigMap as WebsiteConfig;
});
</script>

<style scoped lang="scss">
.footer {
  padding: 0 5px;

  .el-form-item {
    margin-bottom: 0;
  }
}
</style>
