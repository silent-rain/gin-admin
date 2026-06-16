<script setup lang="ts">
import ApiHttpPermissionTab from './ApiHttpPermissionTab.vue'
import MenuPermissionTab from './MenuPermissionTab.vue'

const props = withDefaults(
  defineProps<{
    modelValue: boolean
    roleId: number
  }>(),
  {},
)

const emit = defineEmits(['update:modelValue'])

const activeName = ref('menu')

function handleClose(done: () => void) {
  emit('update:modelValue', false)
  done()
}
</script>

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
            <MenuPermissionTab :role-id="props.roleId" />
          </el-tab-pane>
          <el-tab-pane label="接口权限" name="apiHttp">
            <ApiHttpPermissionTab
              :role-id="props.roleId"
            />
          </el-tab-pane>
        </el-tabs>
      </el-scrollbar>
    </div>
  </el-drawer>
</template>

<style scoped lang="scss">
.details {
  height: 80vh;
}
</style>
