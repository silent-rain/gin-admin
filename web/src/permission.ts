/*
 * @Author: silent-rain
 * @Date: 2023-01-06 23:20:53
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-15 02:07:40
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/permission.ts
 * @Descripttion: 权限守护路由
 */
import router from '@/router';
import {
  filterAsyncRouter,
  progressClose,
  progressStart,
} from '@/hooks/use-permission';
import { useBasicStore } from '@/store/basic';
import { userInfo } from '@/api/user';
import { langTitle } from '@/hooks/use-common';

// no redirect whitelist
const whiteList = ['/login', '/404', '/401'];

// 路由进入前拦截
// to:将要进入的页面 vue-router4.0 不推荐使用next()
router.beforeEach(async (to) => {
  progressStart();
  // i18 page title
  document.title = langTitle(to.meta?.title);
  const basicStore = useBasicStore();

  // 1. 判断 Token
  if (!basicStore.token) {
    if (!whiteList.includes(to.path)) {
      return `/login?redirect=${to.path}`;
    }
    return true;
  }

  // 2. 存在 Token 跳转至首页
  if (to.path === '/login') {
    return '/';
  }

  // 3.判断是否获取用户信息
  if (basicStore.getUserInfo) {
    return true;
  }

  try {
    // 4. 获取用户信息
    const userData = await userInfo({});
    // 5. 动态路由权限筛选
    filterAsyncRouter(userData);
    // 6. 保存用户信息到 store
    basicStore.setUserInfo(userData);
    // 7. 再次执行路由跳转
    return { ...to, replace: true };
  } catch (e) {
    console.error(`route permission error ${e}`);
    basicStore.resetState();
    progressClose();
    return `/login?redirect=${to.path}`;
  }
});

// 路由进入后拦截
router.afterEach(() => {
  progressClose();
});
