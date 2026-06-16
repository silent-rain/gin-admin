import router, { catchRoutes, constantRoutes } from '@/router';
import Layout from '@/layout/index.vue';
import { usePermissionStore } from '@/store/permission';
import { useUserStore } from '@/store/user';
export const asyncRoutesByMenus = (menus) => {
    const routes = [];
    const modules = import.meta.glob('../views/**/**.vue');
    for (const menu of menus) {
        const parentNode = {};
        parentNode.path = menu.path;
        parentNode.name = menu.name;
        if (menu.redirect === '' && menu.children.length > 0) {
            parentNode.redirect = menu.children[0].path;
        }
        else {
            parentNode.redirect = menu.redirect;
        }
        if (menu.component === 'Layout') {
            parentNode.component = shallowRef(Layout);
        }
        else {
            const url = menu.component.replace('@', '..');
            parentNode.component = modules[url];
        }
        parentNode.meta = {
            title: menu.title,
            elSvgIcon: menu.el_svg_icon,
            icon: menu.icon,
        };
        parentNode.alwaysShow = menu.always_show === 1;
        parentNode.hidden = menu.hidden === 1;
        if (menu.children) {
            parentNode.children = asyncRoutesByMenus(menu.children);
        }
        routes.push(parentNode);
    }
    return routes;
};
export const filterAsyncRouter = (menus) => {
    const permissionStore = usePermissionStore();
    const asyncRoutes = asyncRoutesByMenus(menus);
    const allRoutes = constantRoutes.concat(asyncRoutes).concat(catchRoutes);
    permissionStore.setFilterAsyncRoutes(menus, asyncRoutes, allRoutes);
    permissionStore.asyncRoutes.forEach((feItem) => router.addRoute(feItem));
};
export const buttonPermissions = (permissions) => {
    const permissionHash = {};
    for (const item of permissions) {
        permissionHash[item.permission] = item.disabled !== 0;
    }
    usePermissionStore().setButtonPermission(permissions, permissionHash);
};
export const hasButtonPermission = (value) => {
    if (usePermissionStore().permissionHash[value] !== undefined) {
        return true;
    }
    return false;
};
export const isDisabledButton = (value) => {
    const perm = usePermissionStore().permissionHash[value];
    if (perm === undefined) {
        return true;
    }
    return perm;
};
export function resetRouter() {
    const routeNameSet = new Set();
    router.getRoutes().forEach((fItem) => {
        if (fItem.name)
            routeNameSet.add(fItem.name);
    });
    routeNameSet.forEach((setItem) => router.removeRoute(setItem));
    usePermissionStore().asyncRoutes.forEach((feItem) => router.addRoute(feItem));
}
export function resetState() {
    resetRouter();
    useUserStore().resetState();
    usePermissionStore().resetState();
}
export function freshRouter(data) {
    resetRouter();
    filterAsyncRouter(data);
}
