<script setup lang="ts">
import { ElMessage } from 'element-plus'
import screenfull from 'screenfull'
// @ts-ignore
import { onMounted, onUnmounted, reactive, toRefs } from 'vue'
import SvgIcon from '@/icons/SvgIcon.vue'

const state = reactive({
  isFullscreen: false,
})
onMounted(() => {
  init()
})
onUnmounted(() => {
  destroy()
})
function toggleScreen() {
  if (!screenfull.isEnabled) {
    ElMessage({
      message: 'you browser can not work',
      type: 'warning',
    })
    return false
  }
  screenfull.toggle()
}
function change() {
  state.isFullscreen = screenfull.isFullscreen
}
function init() {
  if (screenfull.isEnabled) {
    screenfull.on('change', change)
  }
}
function destroy() {
  if (screenfull.isEnabled) {
    screenfull.off('change', change)
  }
}
const { isFullscreen } = toRefs(state)
</script>

<template>
  <SvgIcon
    :icon-class="isFullscreen ? 'exit-fullscreen' : 'fullscreen'"
    style="width: 17px; height: 17px"
    class="mr-12px"
    @click="toggleScreen"
  />
</template>

<style lang="scss" scoped>
.nav-svg-icon {
  font-size: 18px;
  color: #5a5e66;
  margin-top: 4px;
}
</style>
