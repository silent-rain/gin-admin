declare const BasicDemo: {
    path: string;
    component: import("vue").DefineComponent<{}, {}, any>;
    meta: {
        title: string;
        icon: string;
    };
    alwaysShow: boolean;
    children: ({
        path: string;
        component: () => Promise<typeof import("*.vue")>;
        name: string;
        meta: {
            title: string;
            cacheGroup?: undefined;
            activeMenu?: undefined;
            cachePage?: undefined;
            closeTabRmCache?: undefined;
            leaveRmCachePage?: undefined;
        };
        hidden?: undefined;
        alwaysShow?: undefined;
        children?: undefined;
    } | {
        path: string;
        component: () => Promise<typeof import("*.vue")>;
        name: string;
        meta: {
            title: string;
            cacheGroup: string[];
            activeMenu?: undefined;
            cachePage?: undefined;
            closeTabRmCache?: undefined;
            leaveRmCachePage?: undefined;
        };
        hidden?: undefined;
        alwaysShow?: undefined;
        children?: undefined;
    } | {
        path: string;
        name: string;
        hidden: boolean;
        component: () => Promise<typeof import("*.vue")>;
        meta: {
            title: string;
            activeMenu: string;
            cacheGroup?: undefined;
            cachePage?: undefined;
            closeTabRmCache?: undefined;
            leaveRmCachePage?: undefined;
        };
        alwaysShow?: undefined;
        children?: undefined;
    } | {
        path: string;
        component: () => Promise<typeof import("*.vue")>;
        name: string;
        meta: {
            title: string;
            cachePage: boolean;
            closeTabRmCache: boolean;
            cacheGroup?: undefined;
            activeMenu?: undefined;
            leaveRmCachePage?: undefined;
        };
        hidden?: undefined;
        alwaysShow?: undefined;
        children?: undefined;
    } | {
        path: string;
        name: string;
        component: () => Promise<typeof import("*.vue")>;
        meta: {
            title: string;
            cachePage: boolean;
            leaveRmCachePage: boolean;
            cacheGroup?: undefined;
            activeMenu?: undefined;
            closeTabRmCache?: undefined;
        };
        alwaysShow: boolean;
        children: {
            path: string;
            name: string;
            component: () => Promise<typeof import("*.vue")>;
            meta: {
                title: string;
                cachePage: boolean;
                leaveRmCachePage: boolean;
            };
        }[];
        hidden?: undefined;
    })[];
};
export default BasicDemo;
