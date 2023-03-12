<template>
  <el-container :class="classObj">
    <el-aside>
      <Sidebar v-if="settings.showLeftMenu" />
    </el-aside>

    <el-container>
      <el-header :class="headerClassObj">
        <Navbar v-if="settings.showTopNavbar" />
        <TagsView v-if="settings.showTagsView" />
      </el-header>
      <el-main>
        <div
          v-if="basicStore.isMobile() && sidebar.opened"
          class="drawer-bg"
          @click="handleClickOutside"
        />
        <AppMain />
      </el-main>
      <el-footer v-if="!basicStore.isMobile()">
        <Footer v-if="settings.showFooter" />
      </el-footer>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { computed, onBeforeMount } from 'vue';
import Sidebar from './sidebar/index.vue';
import Navbar from './navbar/index.vue';
import TagsView from './tags-view/index.vue';
import AppMain from './app-main/index.vue';
import Footer from './footer/index.vue';
import { useBasicStore } from '@/store/basic';
import { resizeHandler } from '@/hooks/use-layout';

const { sidebar, settings } = useBasicStore();
const basicStore = useBasicStore();

const classObj = computed(() => {
  return {
    container: true,
    'close-sidebar': !sidebar.opened,
    'hide-sidebar': !settings.showLeftMenu,
    'fixed-header': settings.fixedHeader,
    mobile: basicStore.device === 'mobile',
  };
});
const headerClassObj = computed(() => {
  return {
    'show-navbar': settings.showTopNavbar,
    'show-tags-view': settings.showTagsView,
  };
});

onBeforeMount(() => {
  resizeHandler();
});

// 移动端点击隐藏侧边栏
const handleClickOutside = () => {
  basicStore.setSidebarOpen(false);
};
</script>

<style lang="scss" scoped>
.container {
  .el-aside {
    width: var(--side-bar-width);
    background-color: var(--el-menu-bg-color);
    position: fixed;
  }

  .el-container {
    margin-left: var(--side-bar-width);
  }
}
.container.close-sidebar {
  .el-aside {
    width: var(--side-bar-min-width) !important;
  }

  .el-container {
    margin-left: var(--side-bar-min-width);
  }
}
.container.hide-sidebar .el-aside {
  width: 0 !important;

  .el-container {
    margin-left: 0;
  }
}

// header
.container .el-header {
  height: auto;
}
// 固定 header
.container.fixed-header .el-main {
  height: calc(
    100vh - #{var(--nav-bar-height)} - #{var(--tag-view-height)} - #{var(
        --footer-height
      )}
  );
}

.el-footer {
  height: var(--footer-height);
  // position: absolute;
  padding-top: 20px;
  bottom: 0;
}

// 移动端布局
.container.mobile {
  .el-aside {
    z-index: 300;
  }
  .el-container {
    margin-left: var(--side-bar-min-width);
    z-index: 100;
  }
  :not(.close-sidebar) .el-container {
    margin-left: 0;
  }
}

// 移动端主内容区域
.container.mobile.fixed-header .el-main {
  height: calc(100vh - #{var(--nav-bar-height)});
}

// 移动端点击隐藏侧边栏
.drawer-bg {
  background: #000;
  opacity: 0.3;
  width: 100%;
  top: 0;
  height: 100%;
  position: absolute;
  z-index: 999;
}
</style>
