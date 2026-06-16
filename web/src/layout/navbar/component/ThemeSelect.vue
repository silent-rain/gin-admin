<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { reactive, toRefs } from 'vue'
import SvgIcon from '@/icons/SvgIcon.vue'
import { useConfigStore } from '@/store/config'

const configStore = useConfigStore()

const { theme } = storeToRefs(configStore)
const state = reactive({
  themeOptions: [
    { label: 'base', value: 'base-theme' },
    { label: 'dark', value: 'dark' },
    { label: 'lighting', value: 'lighting-theme' },
    { label: 'china-red', value: 'china-red' },
  ],
})
function handleSetTheme(theme) {
  configStore.setTheme(theme)
}
const { themeOptions } = toRefs(state)
</script>

<template>
  <el-dropdown trigger="click" type="primary" @command="handleSetTheme">
    <SvgIcon
      icon-class="theme-icon"
      style="width: 22px; height: 23px"
      class="mr-12px"
    />
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="item in themeOptions"
          :key="item.value"
          :command="item.value"
          :disabled="theme === item.value"
        >
          <h3 class="pt-6px pb-10px text-16px">
            {{ item.label }}
          </h3>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<style scoped lang="scss">
.theme-icon-style {
  height: 23px;
  width: 23px;
  margin-left: 8px;
  margin-right: 8px;
  color: #494949; //#ff9901;
  position: relative;
  font-weight: bold;
  top: 1px;
}
</style>
