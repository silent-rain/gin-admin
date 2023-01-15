<template>
  <svg-icon
    :icon-class="isFullscreen ? 'exit-fullscreen' : 'fullscreen'"
    style="width: 17px; height: 17px"
    class="mr-12px"
    @click="toggleScreen"
  />
</template>

<script setup lang="ts">
// @ts-ignore
import { onMounted, onUnmounted, reactive, toRefs } from 'vue';
import screenfull from 'screenfull';
import { ElMessage } from 'element-plus';
import SvgIcon from '@/icons/SvgIcon.vue';

// 屏幕的状态
const state = reactive({
  isFullscreen: false,
});

onMounted(() => {
  init();
});
onUnmounted(() => {
  destroy();
});

const toggleScreen = () => {
  if (!screenfull.isEnabled) {
    ElMessage({
      message: 'you browser can not work',
      type: 'warning',
    });
    return false;
  }
  // 切换全屏状态
  screenfull.toggle();
};

// 初始化注册screenFull的change事件
const init = () => {
  // 判断是否支持全屏
  if (screenfull.isEnabled) {
    // 开启监听change事件
    screenfull.on('change', change);
  }
};
// 更改当前屏幕的状态
const change = () => {
  // 更新全屏状态
  state.isFullscreen = screenfull.isFullscreen;
};
// 最后注销监听事件
const destroy = () => {
  if (screenfull.isEnabled) {
    screenfull.off('change', change);
  }
};
const { isFullscreen } = toRefs(state);
</script>

<style lang="scss" scoped>
.nav-svg-icon {
  font-size: 18px;
  color: #5a5e66;
  margin-top: 4px;
}
</style>
