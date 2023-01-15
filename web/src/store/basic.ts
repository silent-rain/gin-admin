import { defineStore } from 'pinia';
import type { RouterTypes } from '~/basic';
import defaultSettings from '@/settings';
import { constantRoutes } from '@/router';

export const useBasicStore = defineStore('basic', {
  state: () => {
    return {
      // router
      allRoutes: [] as RouterTypes,
      buttonCodes: [],
      filterAsyncRoutes: [] as RouterTypes,
      // keep-alive
      cachedViews: [] as Array<string>,
      cachedViewsDeep: [] as Array<string>,
      // other
      sidebar: { opened: true },
      // axios req collection
      axiosPromiseArr: [] as Array<ObjKeys>,
      settings: defaultSettings,
    };
  },
  persist: {
    storage: localStorage,
    paths: ['token'],
  },
  actions: {
    // 设置过滤的异步路由
    setFilterAsyncRoutes(routes: RouterTypes) {
      this.$patch((state) => {
        state.filterAsyncRoutes = routes;
        state.allRoutes = constantRoutes.concat(routes);
      });
    },

    // 设置侧边栏，显示/隐藏
    setSidebarOpen(data: any) {
      this.$patch((state) => {
        state.sidebar.opened = data;
      });
    },
    // 点击侧边栏按钮，显示/隐藏
    setToggleSideBar() {
      this.$patch((state) => {
        state.sidebar.opened = !state.sidebar.opened;
      });
    },

    /* keepAlive缓存 */
    addCachedView(view) {
      this.$patch((state) => {
        if (state.cachedViews.includes(view)) return;
        state.cachedViews.push(view);
      });
    },

    delCachedView(view) {
      this.$patch((state) => {
        const index = state.cachedViews.indexOf(view);
        index > -1 && state.cachedViews.splice(index, 1);
      });
    },
    /* third  keepAlive */
    addCachedViewDeep(view) {
      this.$patch((state) => {
        if (state.cachedViewsDeep.includes(view)) return;
        state.cachedViewsDeep.push(view);
      });
    },
    delCacheViewDeep(view) {
      this.$patch((state) => {
        const index = state.cachedViewsDeep.indexOf(view);
        index > -1 && state.cachedViewsDeep.splice(index, 1);
      });
    },
  },
});
