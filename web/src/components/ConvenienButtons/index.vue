<template>
  <div class="convenient-buttons">
    <el-button
      v-if="buttonDict.add"
      type="primary"
      :icon="Plus"
      @click="handleAddEvent"
      >添加
    </el-button>
    <el-popconfirm
      v-if="buttonDict.batchDelete"
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
import { onBeforeMount, ref } from 'vue';

const props = withDefaults(
  defineProps<{
    buttonList?: string[]; // 按钮名称列表; add/batchDelete
  }>(),
  {
    buttonList: () => ['add', 'batchDelete'],
  },
);

const emit = defineEmits(['addEvent', 'batchDeleteEvent']);
const buttonDict = ref<any>({});

onBeforeMount(() => {
  if (props.buttonList.length > 0) {
    const m = {};
    props.buttonList.forEach((v) => {
      m[v] = true;
    });
    buttonDict.value = m;
  }
});

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
