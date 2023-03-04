<template>
  <div id="Sidebar" class="reset-menu-style">
    <div
      v-if="device === 'mobile' && sidebar.opened"
      class="drawer-bg"
      @click="handleClickOutside"
    />

    <!--logo-->
    <Logo v-if="settings.sidebarLogo" :collapse="!sidebar.opened" />

    <!--router menu-->
    <el-scrollbar>
      <el-menu
        class="el-menu-vertical"
        :collapse="!sidebar.opened"
        :default-active="activeMenu"
        :collapse-transition="false"
        mode="vertical"
      >
        <sidebar-item
          v-for="route in allRoutes"
          :key="route.path"
          :item="route"
          :base-path="route.path"
        />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { storeToRefs } from 'pinia/dist/pinia';
import { useRoute } from 'vue-router';
import Logo from './Logo.vue';
import SidebarItem from './SidebarItem.vue';
import { useBasicStore } from '@/store/basic';
import { usePermissionStore } from '@/store/permission';

const { settings, sidebar, device } = storeToRefs(useBasicStore());
const { allRoutes } = storeToRefs(usePermissionStore());
const route = useRoute();
const basicStore = useBasicStore();

const activeMenu = computed(() => {
  const { meta, path } = route;
  // if set path, the sidebar will highlight the path you set
  if (meta.activeMenu) {
    return meta.activeMenu;
  }
  return path;
});

// 隐藏侧边栏
const handleClickOutside = () => {
  basicStore.setSidebarOpen(false);
};
</script>

<style lang="scss" scoped>
// fix open the item style issue
// .el-menu-vertical {
// width: var(--side-bar-width);
// }

.el-scrollbar {
  height: calc(100vh - #{var(--sidebar-logo-container-height)}) !important;
}

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
