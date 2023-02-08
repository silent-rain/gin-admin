import { defineStore } from 'pinia';
import { RouteRawConfig, RouterTypes } from '~/basic';
import { Menu } from '~/api/system/menu';
import { asyncRoutesByMenus } from '@/hooks/use-menu';
import { constantRoutes, catchRoutes } from '@/router';

export const useMenuStore = defineStore('menu', {
  state: () => {
    return {
      // 菜单路由列表
      menus: [] as Menu[],
      // 异步路由列表
      asyncRoutes: [] as RouteRawConfig[],
      // 所有路由
      allRoutes: [] as RouterTypes,
    };
  },
  actions: {
    // 设置异步路由
    setAsyncRoutes(menus: Menu[]) {
      const asyncRoutes = asyncRoutesByMenus(menus);
      const accessRoutes = constantRoutes
        .concat(asyncRoutes)
        .concat(catchRoutes);
      this.$patch((state) => {
        state.menus = menus;
        state.asyncRoutes = asyncRoutes;
        state.allRoutes = accessRoutes;
      });
    },
  },
});
