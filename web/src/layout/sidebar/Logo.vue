<template>
  <div class="sidebar-logo-container" :class="{ collapse: collapse }">
    <transition name="sidebar-logo-fade">
      <!--  折叠显示   -->
      <router-link v-if="collapse" class="sidebar-logo-link" to="/">
        <svg-icon v-if="logo" :icon-class="logo" class="sidebar-logo" />
        <h1 v-else class="sidebar-title">
          {{ title }}
        </h1>
      </router-link>
      <!--  正常显示   -->
      <router-link v-else class="sidebar-logo-link" to="/">
        <svg-icon v-if="logo" :icon-class="logo" class="sidebar-logo" />
        <h1 class="sidebar-title">
          {{ title }}
        </h1>
      </router-link>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { useBasicStore } from '@/store/basic';
import SvgIcon from '@/icons/SvgIcon.vue';

const { settings, webSiteConfigMap } = useBasicStore();
defineProps({
  // 是否折叠
  collapse: {
    type: Boolean,
    required: true,
  },
});

// 获取动态标题
const getTitle = computed(() => {
  let title = webSiteConfigMap['website_title']?.value;
  if (title) {
    return title;
  }
  return settings.title;
});

const state = reactive({
  title: getTitle.value,
  // src/icons/common/sidebar-logo.svg
  logo: 'sidebar-logo',
});

// export to page for use
const { title, logo } = toRefs(state);
</script>

<style lang="scss" scoped>
//vue3.0 过度效果更改  enter-> enter-from   leave-> leave-from
.sidebar-logo-container {
  position: relative;
  width: 100%;
  height: var(--sidebar-logo-container-height);
  line-height: var(--sidebar-logo-container-height);
  background: var(--sidebar-logo-background);
  padding-left: 14px;
  text-align: left;
  overflow: hidden;
  & .sidebar-logo-link {
    height: 100%;
    width: 100%;
    & .sidebar-logo {
      fill: currentColor;
      color: var(--sidebar-logo-color);
      width: var(--sidebar-logo-width);
      height: var(--sidebar-logo-height);
      vertical-align: middle;
      margin-right: 12px;
    }
    & .sidebar-title {
      display: inline-block;
      margin: 0;
      color: var(--sidebar-logo-title-color);
      font-weight: 600;
      line-height: 50px;
      font-size: 14px;
      font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
      vertical-align: middle;
    }
  }
  &.collapse {
    .sidebar-logo {
      margin-right: 0;
    }
  }
}
</style>
