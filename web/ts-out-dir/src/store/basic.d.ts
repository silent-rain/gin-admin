export declare const useBasicStore: import("pinia").StoreDefinition<"basic", {
    cachedViews: Array<string>;
    cachedViewsDeep: Array<string>;
    sidebar: {
        opened: boolean;
    };
    axiosPromiseArr: Array<ObjKeys>;
    settings: import("../typings/settings").SettingsConfig;
    osType: number;
    device: string;
    webSiteConfigMap: {};
}, {}, {
    setSidebarOpen(data: any): void;
    setToggleSideBar(): void;
    setOsType(): void;
    setDevice(): void;
    isMobile(): boolean;
    setWebSiteConfig(data: any): void;
    addCachedView(view: any): void;
    delCachedView(view: any): void;
    addCachedViewDeep(view: any): void;
    delCacheViewDeep(view: any): void;
}>;
