<template>
  <el-tabs v-model="activeName">
    <el-tab-pane label="网站设置" name="websiteConfig">
      <WebsiteConfig></WebsiteConfig>
    </el-tab-pane>
  </el-tabs>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import WebsiteConfig from './components/WebsiteConfig.vue';

const activeName = ref('websiteConfig');
</script>

<style scoped lang="scss"></style>
