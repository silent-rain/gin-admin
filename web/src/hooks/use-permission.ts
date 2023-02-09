import { RouteRecordName } from 'vue-router';
import router, { catchRoutes, constantRoutes } from '@/router';
import Layout from '@/layout/index.vue';
import { usePermissionStore } from '@/store/permission';
import { RouteRawConfig } from '~/basic';
import { Menu } from '~/api/system/menu';

// 菜单列表转为路由列表
export const asyncRoutesByMenus = (menus: Menu[]) => {
  const routes: RouteRawConfig[] = [];
  for (const menu of menus) {
    const parentNode: RouteRawConfig = {} as RouteRawConfig;
    parentNode.path = menu.path;
    if (menu.component === 'Layout') {
      parentNode.component = shallowRef(Layout);
    } else {
      const url = menu.component.replace('@', '..');
      console.log(url);
      // parentNode.component = importModule(url);
      parentNode.component = () =>
        import(/* @vite-ignore */ '@/views/system/menu/index.vue');
      // parentNode.component = defineAsyncComponent(
      //   () => import(/* @vite-ignore */ url),
      // );
    }
    parentNode.meta = { title: menu.title, elSvgIcon: menu.icon };
    // parentNode.alwaysShow = menu.always_show;
    parentNode.hidden = menu.hidden === 1 ? true : false;

    if (menu.children) {
      parentNode.children = asyncRoutesByMenus(menu.children);
    }
    routes.push(parentNode);
  }
  return routes;
};

// 动态加载模块
function importModule(path: string) {
  // who knows what will be imported here?
  return defineAsyncComponent(() => import(/* @vite-ignore */ path));
}

// 过滤异步路由
export const filterAsyncRouter = (menus: Menu[]) => {
  const asyncRoutes = asyncRoutesByMenus(menus);
  const allRoutes = constantRoutes.concat(asyncRoutes).concat(catchRoutes);
  usePermissionStore().setFilterAsyncRoutes(menus, asyncRoutes, allRoutes);
  asyncRoutes.forEach((route) => {
    router.addRoute(route);
  });
  console.log(router.getRoutes());
};

//重置路由
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
  usePermissionStore().resetState();
}

//刷新路由
export function freshRouter(data: any) {
  resetRouter();
  filterAsyncRouter(data);
  // location.reload()
}
