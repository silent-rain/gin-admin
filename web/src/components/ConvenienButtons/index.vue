<template>
  <div class="convenient-buttons">
    <el-button type="primary" :icon="Plus" @click="handleAddEvent"
      >添加
    </el-button>
    <el-popconfirm
      confirm-button-text="确认"
      cancel-button-text="取消"
      :icon="InfoFilled"
      icon-color="#E6A23C"
      title="确定删除吗?"
      @confirm="handleBatchDeleteEvent"
      @cancel="handleBatchDeleteCancelEvent"
    >
      <template #reference>
        <el-button type="danger" :icon="Delete">批量删除 </el-button>
      </template>
    </el-popconfirm>
  </div>
</template>

<script setup lang="ts">
import { Delete, Plus, InfoFilled } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import { onBeforeMount } from 'vue';

const emit = defineEmits(['addEvent', 'batchDeleteEvent']);

onBeforeMount(() => {});

// 添加事件
const handleAddEvent = () => {
  emit('addEvent', null);
};

// 批量删除事件
const handleBatchDeleteEvent = () => {
  emit('batchDeleteEvent', null);
};

// 批量删除取消事件
const handleBatchDeleteCancelEvent = () => {
  ElMessage.warning('取消操作');
};
</script>

<style scoped lang="scss"></style>
