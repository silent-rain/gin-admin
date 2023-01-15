/*
 * @Author: silent-rain
 * @Date: 2023-01-06 23:20:53
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-15 02:07:06
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/hooks/use-self-router.ts
 * @Descripttion:
 */
import router from '@/router';

export const getQueryParam = () => {
  const route: any = router.currentRoute;
  if (route.value?.query.params) {
    return JSON.parse(route.value.query.params);
  }
};
// vue router
export const routerPush = (name, params) => {
  let data = {};
  if (params) {
    data = {
      params: JSON.stringify(params),
    };
  } else {
    data = {};
  }
  router.push({
    name,
    query: data,
  });
};
export const routerReplace = (name, params) => {
  let data = {};
  if (params) {
    data = {
      params: JSON.stringify(params),
    };
  } else {
    data = {};
  }
  router.replace({
    name,
    query: data,
  });
};

export const routeInfo = () => {
  return router.currentRoute;
};
export const routerBack = () => {
  router.go(-1);
};
