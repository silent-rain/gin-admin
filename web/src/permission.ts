import router from '@/router';
import { usePermissionStore } from '@/store/permission';
import { progressClose, progressStart } from '@/hooks/use-basic';
import { langTitle } from '@/hooks/use-common';
import { buttonPermissions, filterAsyncRouter } from './hooks/use-permission';
import { useUserStore } from '@/store/user';
import { useBasicStore } from '@/store/basic';

// no redirect whitelist
const whiteList = ['/login', '/register', '/404', '/401'];

// 路由进入前拦截
// to:将要进入的页面 vue-router4.0 不推荐使用next()
router.beforeEach(async (to, from) => {
  progressStart();

  // i18 page title
  document.title = langTitle(to.meta?.title);
  const userStore = useUserStore();
  const permissionStore = usePermissionStore();
  const basicStore = useBasicStore();

  // 设置 Device
  basicStore.setDevice();

  // 判断 Token, 不存在则跳转至登录
  if (!userStore.token) {
    if (!whiteList.includes(to.path)) {
      return `/login?redirect=${to.path}`;
    }
    return true;
  }

  // 存在 Token 跳转至首页
  if (to.path === '/login') {
    return '/';
  }

  // 判断是否获取用户信息
  if (userStore.getUserInfo && permissionStore.allRoutes.length > 0) {
    return true;
  }

  try {
    // 获取用户信息
    const userData = await userStore.userInfo();
    // 设置用户信息
    userStore.setUserInfo(userData);
    // 设置动态路由权限
    filterAsyncRouter(userData.menus);
    // 设置按钮权限
    buttonPermissions(userData.permissions);

    // 执行路由跳转
    return { ...to, replace: true };
  } catch (e) {
    console.error(`route permission error ${e}`);
    userStore.resetState();
    progressClose();
    return `/login?redirect=${to.path}`;
  }
});

// 路由进入后拦截
router.afterEach(() => {
  progressClose();
});
