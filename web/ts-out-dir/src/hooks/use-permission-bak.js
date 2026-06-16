import Layout from '@/layout/index.vue';
import 'nprogress/nprogress.css';
const buttonCodes = [];
export const filterAsyncRoutesByMenuList = (menuList) => {
    const filterRouter = [];
    menuList.forEach((route) => {
        if (route.category === 3) {
            buttonCodes.push(route.code);
        }
        else {
            const itemFromReqRouter = getRouteItemFromReqRouter(route);
            if (route.children?.length) {
                itemFromReqRouter.children = filterAsyncRoutesByMenuList(route.children);
            }
            filterRouter.push(itemFromReqRouter);
        }
    });
    return filterRouter;
};
const getRouteItemFromReqRouter = (route) => {
    const tmp = { meta: { title: '' } };
    const routeKeyArr = [
        'path',
        'component',
        'redirect',
        'alwaysShow',
        'name',
        'hidden',
    ];
    const metaKeyArr = ['title', 'activeMenu', 'elSvgIcon', 'icon'];
    const modules = import.meta.glob('../views/**/**.vue');
    routeKeyArr.forEach((fItem) => {
        if (fItem === 'component') {
            if (route[fItem] === 'Layout') {
                tmp[fItem] = Layout;
            }
            else {
                tmp[fItem] = modules[`../views/${route[fItem]}`];
            }
        }
        else if (fItem === 'path' && route.parentId === 0) {
            tmp[fItem] = `/${route[fItem]}`;
        }
        else if (['hidden', 'alwaysShow'].includes(fItem)) {
            tmp[fItem] = !!route[fItem];
        }
        else if (['name'].includes(fItem)) {
            tmp[fItem] = route.code;
        }
        else if (route[fItem]) {
            tmp[fItem] = route[fItem];
        }
    });
    metaKeyArr.forEach((fItem) => {
        if (route[fItem] && tmp.meta)
            tmp.meta[fItem] = route[fItem];
    });
    if (route.extra) {
        Object.entries(route.extra.parse(route.extra)).forEach(([key, value]) => {
            if (key === 'meta' && tmp.meta) {
                tmp.meta[key] = value;
            }
            else {
                tmp[key] = value;
            }
        });
    }
    return tmp;
};
export function filterAsyncRoutesByRoles(routes, roles) {
    const res = [];
    routes.forEach((route) => {
        const tmp = { ...route };
        if (hasPermission(roles, tmp)) {
            if (tmp.children) {
                tmp.children = filterAsyncRoutesByRoles(tmp.children, roles);
            }
            res.push(tmp);
        }
    });
    return res;
}
function hasPermission(roles, route) {
    if (route?.meta?.roles) {
        return roles?.some((role) => route.meta?.roles?.includes(role));
    }
    return true;
}
export function filterAsyncRouterByCodes(codesRoutes, codes) {
    const filterRouter = [];
    codesRoutes.forEach((routeItem) => {
        if (hasCodePermission(codes, routeItem)) {
            if (routeItem.children)
                routeItem.children = filterAsyncRouterByCodes(routeItem.children, codes);
            filterRouter.push(routeItem);
        }
    });
    return filterRouter;
}
function hasCodePermission(codes, routeItem) {
    if (routeItem.meta?.code) {
        return codes.includes(routeItem.meta.code) || routeItem.hidden;
    }
    return true;
}
