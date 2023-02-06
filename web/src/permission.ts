import router from '@/router';
import {
  filterAsyncRouter,
  progressClose,
  progressStart,
} from '@/hooks/use-permission';
import { useUserStore } from '@/store/user';
import { langTitle } from '@/hooks/use-common';

// no redirect whitelist
const whiteList = ['/login', '/register', '/404', '/401'];

// 路由进入前拦截
// to:将要进入的页面 vue-router4.0 不推荐使用next()
router.beforeEach(async (to) => {
  progressStart();
  // i18 page title
  document.title = langTitle(to.meta?.title);
  const userStore = useUserStore();

  // 判断 Token
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
  if (userStore.getUserInfo) {
    return true;
  }

  try {
    // 获取用户信息
    const userData = await userStore.userInfo();
    // 保存用户信息到 store
    userStore.setUserInfo(userData);
    // 动态路由权限筛选
    filterAsyncRouter(userData);
    // 再次执行路由跳转
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
