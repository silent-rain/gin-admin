import router from '@/router';
import { usePermissionStore } from '@/store/permission';
import { progressClose, progressStart, fecthWebSiteConfigList, } from '@/hooks/use-basic';
import { langTitle } from '@/hooks/use-common';
import { buttonPermissions, filterAsyncRouter } from './hooks/use-permission';
import { useUserStore } from '@/store/user';
import { useBasicStore } from '@/store/basic';
const whiteList = ['/login', '/register', '/404', '/401'];
router.beforeEach(async (to, from) => {
    progressStart();
    document.title = langTitle(to.meta?.title);
    const userStore = useUserStore();
    const permissionStore = usePermissionStore();
    const basicStore = useBasicStore();
    basicStore.setDevice();
    if (Object.keys(basicStore.webSiteConfigMap).length === 0) {
        fecthWebSiteConfigList();
    }
    if (!userStore.token) {
        if (!whiteList.includes(to.path)) {
            return `/login?redirect=${to.path}`;
        }
        return true;
    }
    if (to.path === '/login') {
        return '/';
    }
    if (userStore.getUserInfo && permissionStore.allRoutes.length > 0) {
        return true;
    }
    try {
        const userData = await userStore.userInfo();
        userStore.setUserInfo(userData);
        filterAsyncRouter(userData.menus);
        buttonPermissions(userData.permissions);
        return { ...to, replace: true };
    }
    catch (e) {
        console.error(`route permission error ${e}`);
        userStore.resetState();
        progressClose();
        return `/login?redirect=${to.path}`;
    }
});
router.afterEach(() => {
    progressClose();
});
