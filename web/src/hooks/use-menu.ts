import Layout from '@/layout/index.vue';
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
      // parentNode.component = importModule(url);
      parentNode.component = () => import(/* @vite-ignore */ url);
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
