<template>
  <el-container :class="classObj">
    <!-- 侧边栏 -->
    <el-aside v-if="settings.showLeftMenu">
      <Sidebar />
    </el-aside>

    <el-container>
      <!-- 顶部导航栏 -->
      <el-header :class="headerClassObj">
        <Navbar v-if="settings.showTopNavbar" />
        <TagsView v-if="settings.showTagsView" />
      </el-header>
      <!-- 主内容区域 -->
      <el-main>
        <div
          v-if="sidebar.opened"
          class="drawer-bg"
          @click="handleClickOutside"
        />
        <AppMain />
      </el-main>
      <!-- 页脚 -->
      <el-footer v-if="settings.showFooter">
        <Footer />
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
    containerx: true,
    headerx: true,
    'close-sidebar': !sidebar.opened,
    'hide-sidebar': !settings.showLeftMenu,
    'fixed-header': settings.fixedHeader,
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
// pc 默认展开样式
.containerx {
  .el-aside {
    width: var(--side-bar-width);
    background-color: var(--el-menu-bg-color);
    position: fixed;
  }

  .el-container {
    margin-left: var(--side-bar-width);

    // 顶部导航
    .el-header {
      --el-header-padding: 0 0;

      // 侧边栏收缩按钮
      :deep(.hamburger-container) {
        margin-left: 10px;
      }
    }
  }
}

// pc 默收缩样式
.containerx.close-sidebar {
  .el-aside {
    width: var(--side-bar-min-width) !important;
  }

  .el-container {
    margin-left: var(--side-bar-min-width);
  }
}

// pc 隐藏侧边栏
.containerx.hide-sidebar {
  .el-aside {
    width: 0 !important;
  }
  // 右侧层容器
  .el-container {
    margin-left: 0;

    .el-header {
      // 侧边栏收缩按钮
      :deep(.hamburger-container) {
        display: none;
      }
      // 面包屑导航
      // :deep(.el-breadcrumb) {
      //   display: none;
      // }
    }
  }
}

// header
// .containerx .el-header {
//   height: auto;
// }
// 固定 header
// .containerx.fixed-header .el-main {
//   height: calc(
//     100vh - #{var(--nav-bar-height)} - #{var(--tag-view-height)} - #{var(--footer-height)}
//   );
// }
// .el-main {
//   min-height: 100%;
// }

.el-footer {
  height: var(--footer-height);
  width: 100%;
  background: #fafafa;
  // position: absolute;
  bottom: 0;
}

// .el-footer {
//   height: var(--footer-height);
//   background: #fafafa;
//   position: absolute;
//   bottom: 0;
// }

// 移动端布局
@media screen and (max-width: 760px) {
  // 展开样式
  .containerx {
    .el-aside {
      width: var(--side-bar-width);
      z-index: 300;
      background-color: var(--el-menu-bg-color);
      position: fixed;
    }

    .el-container {
      margin-left: 0;

      // 顶部导航
      .el-header {
        // 侧边栏收缩按钮
        :deep(.hamburger-container) {
          z-index: 300;
          margin-left: calc(#{var(--side-bar-width)} + 10px);
          display: block;
        }
        // 面包屑导航
        :deep(.el-breadcrumb) {
          display: none;
        }
        // tag 标签
        :deep(.tags-view-container) {
          display: none;
        }
      }

      .el-main {
        padding: 10px;
      }
    }
  }

  // 收缩样式
  .containerx.close-sidebar {
    .el-aside {
      z-index: 300;
      display: none;
    }
    .el-container {
      margin-left: 0;

      // 顶部导航
      .el-header {
        // 侧边栏收缩按钮
        :deep(.hamburger-container) {
          z-index: 300;
          margin-left: 10px;
          display: block;
        }
      }
    }
  }

  // 移动端主内容区域
  // .container.mobile.fixed-header .el-main {
  //   height: calc(100vh - #{var(--nav-bar-height)});
  // }

  // 移动端点击隐藏侧边栏
  .drawer-bg {
    height: 100%;
    width: 100%;
    position: fixed;
    margin-left: -10px;
    background: #000;
    opacity: 0.3;
    top: 0;
    z-index: 200;
  }
}
</style>
