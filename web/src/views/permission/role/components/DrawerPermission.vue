<template>
  <el-drawer
    :model-value="modelValue"
    title="分配权限"
    :before-close="handleClose"
    size="500px"
  >
    <div class="details">
      <el-scrollbar height="80vh">
        <el-tabs v-model="activeName">
          <el-tab-pane label="菜单权限" name="menu">
            <MenuPermissionTab :roleId="props.roleId"></MenuPermissionTab>
          </el-tab-pane>
          <el-tab-pane label="接口权限" name="apiHttp">
            <ApiHttpPermissionTab :roleId="props.roleId"></ApiHttpPermissionTab>
          </el-tab-pane>
        </el-tabs>
      </el-scrollbar>
    </div>
  </el-drawer>
</template>

<script setup lang="ts">
import MenuPermissionTab from './MenuPermissionTab.vue';
import ApiHttpPermissionTab from './ApiHttpPermissionTab.vue';

const props = withDefaults(
  defineProps<{
    modelValue: boolean;
    roleId: number;
  }>(),
  {},
);

const emit = defineEmits(['update:modelValue']);

const activeName = ref('menu');

const handleClose = (done: () => void) => {
  emit('update:modelValue', false);
  done();
};
</script>

<style scoped lang="scss">
.details {
  height: 80vh;
}
</style>
