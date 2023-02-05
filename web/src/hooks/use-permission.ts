import NProgress from 'nprogress';
import type { RouteRecordName } from 'vue-router';
import type { RouteRawConfig, RouterTypes, rawConfig } from '~/basic';
/**
 * 根据请求，过滤异步路由
 * @param:menuList 异步路由数组
 * return 过滤后的异步路由
 */
import Layout from '@/layout/index.vue';
/*
 * 路由操作
 * */
import router, { asyncRoutes, constantRoutes, roleCodeRoutes } from '@/router';
// 进度条
import 'nprogress/nprogress.css';
import { useBasicStore } from '@/store/basic';
import { useUserStore } from '@/store/user';

const buttonCodes: Array<number> = []; // 按钮权限
interface menuRow {
  category: number;
  code: number;
  children: RouterTypes;
}

// 过滤菜单路由数组
export const filterAsyncRoutesByMenuList = (menuList: menuRow[]) => {
  const filterRouter: RouterTypes = [];
  menuList.forEach((route: menuRow) => {
    // button permission
    if (route.category === 3) {
      buttonCodes.push(route.code);
    } else {
      // generator every router item by menuList
      const itemFromReqRouter = getRouteItemFromReqRouter(route);
      if (route.children?.length) {
        // judge  the type is router or button
        itemFromReqRouter.children = filterAsyncRoutesByMenuList(
          route.children,
        );
      }
      filterRouter.push(itemFromReqRouter);
    }
  });
  return filterRouter;
};

const getRouteItemFromReqRouter = (route): RouteRawConfig => {
  const tmp: rawConfig = { meta: { title: '' } };
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
  // generator routeKey
  routeKeyArr.forEach((fItem: string) => {
    if (fItem === 'component') {
      if (route[fItem] === 'Layout') {
        tmp[fItem] = Layout;
      } else {
        // has error , i will fix it through plugins
        // tmp[fItem] = () => import(`@/views/permission-center/test/TestTableQuery.vue`)
        tmp[fItem] = modules[`../views/${route[fItem]}`];
      }
    } else if (fItem === 'path' && route.parentId === 0) {
      tmp[fItem] = `/${route[fItem]}`;
    } else if (['hidden', 'alwaysShow'].includes(fItem)) {
      tmp[fItem] = !!route[fItem];
    } else if (['name'].includes(fItem)) {
      tmp[fItem] = route.code;
    } else if (route[fItem]) {
      tmp[fItem] = route[fItem];
    }
  });
  // generator metaKey
  metaKeyArr.forEach((fItem) => {
    if (route[fItem] && tmp.meta) tmp.meta[fItem] = route[fItem];
  });
  // route extra insert
  if (route.extra) {
    Object.entries(route.extra.parse(route.extra)).forEach(([key, value]) => {
      if (key === 'meta' && tmp.meta) {
        tmp.meta[key] = value;
      } else {
        tmp[key] = value;
      }
    });
  }
  return tmp as RouteRawConfig;
};

/**
 * 根据角色数组过滤异步路由
 * @param routes asyncRoutes 未过滤的异步路由
 * @param roles  角色数组
 * return 过滤后的异步路由
 */
export function filterAsyncRoutesByRoles(
  routes: RouteRawConfig[],
  roles: number[],
) {
  const res: RouterTypes = [];
  routes.forEach((route) => {
    const tmp: RouteRawConfig = { ...route };
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutesByRoles(tmp.children, roles);
      }
      res.push(tmp);
    }
  });
  return res;
}

// 是否存在 role 权限
function hasPermission(roles: number[], route: RouteRawConfig) {
  if (route?.meta?.roles) {
    return roles?.some((role) => route.meta?.roles?.includes(role));
  }
  return true;
}

/**
 * 根据code数组，过滤异步路由
 * @param codes code数组
 * @param codesRoutes 未过滤的异步路由
 * return 过滤后的异步路由
 */
export function filterAsyncRouterByCodes(
  codesRoutes: RouteRawConfig[],
  codes: number[],
) {
  const filterRouter: RouterTypes = [];
  codesRoutes.forEach((routeItem: RouteRawConfig) => {
    if (hasCodePermission(codes, routeItem)) {
      if (routeItem.children)
        routeItem.children = filterAsyncRouterByCodes(
          routeItem.children,
          codes,
        );
      filterRouter.push(routeItem);
    }
  });
  return filterRouter;
}

// 是否存在 code
function hasCodePermission(codes: number[], routeItem: RouteRawConfig) {
  if (routeItem.meta?.code) {
    return codes.includes(routeItem.meta.code) || routeItem.hidden;
  }
  return true;
}

// 过滤异步路由
export function filterAsyncRouter({ menus, roles, codes }) {
  const basicStore = useBasicStore();
  const permissionMode = basicStore.settings?.permissionMode;
  let accessRoutes: RouterTypes = [];
  if (permissionMode === 'rbac') {
    accessRoutes = filterAsyncRoutesByMenuList(menus); // by menuList
  } else if (permissionMode === 'roles') {
    accessRoutes = filterAsyncRoutesByRoles(roleCodeRoutes, roles); // by roles
  } else {
    accessRoutes = filterAsyncRouterByCodes(roleCodeRoutes, codes); // by codes
  }
  accessRoutes.forEach((route) => router.addRoute(route));
  asyncRoutes.forEach((item) => router.addRoute(item));
  basicStore.setFilterAsyncRoutes(accessRoutes);
}

// 重置路由
export function resetRouter() {
  // 移除之前存在的路由
  const routeNameSet: Set<RouteRecordName> = new Set();
  router.getRoutes().forEach((fItem) => {
    if (fItem.name) routeNameSet.add(fItem.name);
  });
  routeNameSet.forEach((setItem) => router.removeRoute(setItem));
  // 新增constantRoutes
  constantRoutes.forEach((feItem) => router.addRoute(feItem));
}

// 重置登录状态
export function resetState() {
  resetRouter();
  useUserStore().resetState();
}

// 刷新路由
export function freshRouter(data: any) {
  resetRouter();
  filterAsyncRouter(data);
  // location.reload()
}

NProgress.configure({ showSpinner: false });
// 开始进度条
export const progressStart = () => {
  NProgress.start();
};
// 关闭进度条
export const progressClose = () => {
  NProgress.done();
};
