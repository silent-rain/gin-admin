import { defineStore } from 'pinia';
import defaultSettings from '@/settings';
export const useBasicStore = defineStore('basic', {
    state: () => {
        return {
            cachedViews: [],
            cachedViewsDeep: [],
            sidebar: { opened: true },
            axiosPromiseArr: [],
            settings: defaultSettings,
            osType: 0,
            device: 'desktop',
            webSiteConfigMap: {},
        };
    },
    persist: {
        storage: localStorage,
        pick: ['token'],
    },
    actions: {
        setSidebarOpen(data) {
            this.$patch((state) => {
                state.sidebar.opened = data;
            });
        },
        setToggleSideBar() {
            this.$patch((state) => {
                state.sidebar.opened = !state.sidebar.opened;
            });
        },
        setOsType() {
            const ua = window.navigator.userAgent;
            let osType = 0;
            if (/(Android)/.test(ua)) {
                osType = 1;
            }
            else if (/(iPhone|iPad)/.test(ua)) {
                osType = 2;
            }
            else {
                osType = 3;
            }
            this.$patch((state) => {
                state.osType = osType;
            });
        },
        setDevice() {
            const WIDTH = 992;
            const isMobile = document.body.getBoundingClientRect().width - 1 < WIDTH;
            this.$patch((state) => {
                state.device = isMobile ? 'mobile' : 'desktop';
                state.sidebar.opened = !isMobile;
            });
            this.setOsType();
        },
        isMobile() {
            return this.device === 'mobile' ? true : false;
        },
        setWebSiteConfig(data) {
            this.$patch((state) => {
                state.webSiteConfigMap = data;
            });
        },
        addCachedView(view) {
            this.$patch((state) => {
                if (state.cachedViews.includes(view))
                    return;
                state.cachedViews.push(view);
            });
        },
        delCachedView(view) {
            this.$patch((state) => {
                const index = state.cachedViews.indexOf(view);
                index > -1 && state.cachedViews.splice(index, 1);
            });
        },
        addCachedViewDeep(view) {
            this.$patch((state) => {
                if (state.cachedViewsDeep.includes(view))
                    return;
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
