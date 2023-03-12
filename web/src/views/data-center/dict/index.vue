<template>
  <el-row :gutter="20">
    <el-col :span="6">
      <Dict></Dict>
    </el-col>
    <el-col :span="18">
      <DictData></DictData>
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import Dict from './components/Dict.vue';
import DictData from './components/DictData.vue';
</script>

<style scoped lang="scss"></style>
