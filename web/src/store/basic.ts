import { defineStore } from 'pinia';
import defaultSettings from '@/settings';

export const useBasicStore = defineStore('basic', {
  state: () => {
    return {
      buttonCodes: [],
      // keep-alive
      cachedViews: [] as Array<string>,
      cachedViewsDeep: [] as Array<string>,
      // other
      sidebar: { opened: true },
      // axios req collection
      axiosPromiseArr: [] as Array<ObjKeys>,
      settings: defaultSettings,
      // 终端类型: 0:未知,1:安卓,2:ios,3:web
      osType: 0,
    };
  },
  persist: {
    storage: localStorage,
    paths: ['token'],
  },
  actions: {
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

    // 设置终端类型
    setOsType() {
      const ua = window.navigator.userAgent;
      let osType = 0;
      if (/(Android)/.test(ua)) {
        osType = 1;
      } else if (/(iPhone|iPad)/.test(ua)) {
        osType = 2;
      } else {
        osType = 3;
      }
      this.$patch((state) => {
        state.osType = osType;
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
