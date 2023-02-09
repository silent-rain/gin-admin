import { defineStore } from 'pinia';
import { RouteRawConfig, RouterTypes } from '~/store/router';
import { Menu } from '~/api/system/menu';

export const usePermissionStore = defineStore('permission', {
  state: () => {
    return {
      // 菜单路由列表
      menus: [] as Menu[],
      // 按钮权限
      buttonCodes: [] as any[],
      // 异步路由列表
      asyncRoutes: [] as RouteRawConfig[],
      // 所有路由
      allRoutes: [] as RouterTypes,
    };
  },
  actions: {
    // 设置过滤异步路由
    setFilterAsyncRoutes(
      menus: Menu[],
      asyncRoutes: RouteRawConfig[],
      allRoutes: RouteRawConfig[],
    ) {
      this.$patch((state) => {
        state.menus = menus;
        state.asyncRoutes = asyncRoutes;
        state.allRoutes = allRoutes;
      });
    },
    // 重置状态
    resetState() {
      this.$patch((state) => {
        state.menus = [];
        // state.codes = [];
        state.allRoutes = [];
        state.buttonCodes = [];
        state.asyncRoutes = [];
      });
    },
  },
});
