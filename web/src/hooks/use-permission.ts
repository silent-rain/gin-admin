import { RouteRecordName } from 'vue-router';
import router, { catchRoutes, constantRoutes } from '@/router';
import Layout from '@/layout/index.vue';
import { usePermissionStore } from '@/store/permission';
import { useUserStore } from '@/store/user';
import { RouteRawConfig } from '~/store/router';
import { ButtonPermission, Menu } from '~/api/permission/menu';

// 菜单列表转为路由列表
export const asyncRoutesByMenus = (menus: Menu[]) => {
  const routes: RouteRawConfig[] = [];
  const modules = import.meta.glob('../views/**/**.vue');
  for (const menu of menus) {
    const parentNode: RouteRawConfig = {} as RouteRawConfig;
    parentNode.path = menu.path;
    parentNode.name = menu.name;
    // 重定向处理, 父级路由如果没有设置则默认重定向第一个子路由
    if (menu.redirect === '' && menu.children.length > 0) {
      parentNode.redirect = menu.children[0].path;
    } else {
      parentNode.redirect = menu.redirect;
    }

    // 组件处理
    if (menu.component === 'Layout') {
      parentNode.component = shallowRef(Layout);
    } else {
      const url = menu.component.replace('@', '..');
      parentNode.component = modules[url];
    }

    // 元数据处理
    parentNode.meta = {
      title: menu.title,
      elSvgIcon: menu.el_svg_icon,
      icon: menu.icon,
    };
    parentNode.alwaysShow = menu.always_show === 1;
    parentNode.hidden = menu.hidden === 1;

    // 子路由处理
    if (menu.children) {
      parentNode.children = asyncRoutesByMenus(menu.children);
    }
    routes.push(parentNode);
  }
  return routes;
};

// 过滤异步路由
export const filterAsyncRouter = (menus: Menu[]) => {
  const permissionStore = usePermissionStore();
  const asyncRoutes = asyncRoutesByMenus(menus);
  const allRoutes = constantRoutes.concat(asyncRoutes).concat(catchRoutes);
  permissionStore.setFilterAsyncRoutes(menus, asyncRoutes, allRoutes);

  // 新增异步路由
  permissionStore.asyncRoutes.forEach((feItem) => router.addRoute(feItem));
};

// 按钮权限
export const buttonPermissions = (permissions: ButtonPermission[]) => {
  const permissionHash = {};
  for (const item of permissions) {
    permissionHash[item.permission] = item.disabled !== 0;
  }
  usePermissionStore().setButtonPermission(permissions, permissionHash);
};
// 是否存在按钮权限
export const hasButtonPermission = (value: string): boolean => {
  if (usePermissionStore().permissionHash[value] !== undefined) {
    return true;
  }
  return false;
};
// 按钮是否禁用
export const isDisabledButton = (value: string): boolean => {
  const perm = usePermissionStore().permissionHash[value];
  // 不存在按钮权限, 则禁用
  if (perm === undefined) {
    return true;
  }
  return perm;
};

// 重置路由
export function resetRouter() {
  // 移除之前存在的路由
  const routeNameSet: Set<RouteRecordName> = new Set();
  router.getRoutes().forEach((fItem) => {
    if (fItem.name) routeNameSet.add(fItem.name);
  });
  routeNameSet.forEach((setItem) => router.removeRoute(setItem));

  // 新增异步路由
  usePermissionStore().asyncRoutes.forEach((feItem) => router.addRoute(feItem));
}

// 重置登录状态
export function resetState() {
  resetRouter();
  useUserStore().resetState();
  usePermissionStore().resetState();
}

// 刷新路由
export function freshRouter(data: any) {
  resetRouter();
  filterAsyncRouter(data);
  // location.reload()
}
