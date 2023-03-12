<template>
  <el-row :gutter="20">
    <el-col :span="6" :xs="24">
      <Dict v-model:dictId="dictId"></Dict>
    </el-col>
    <el-col :span="18" :xs="24">
      <DictData :dictId="dictId"></DictData>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import Dict from './components/Dict.vue';
import DictData from './components/DictData.vue';

const dictId = ref<number>(0);
</script>

<style scoped lang="scss"></style>
