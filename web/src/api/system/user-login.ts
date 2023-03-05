/**用户登录管理
 *
 */

import axiosReq from '@/utils/axios-req';

// 获取用户登录信息列表
export const getUserLoginList = async (params: any) => {
  return axiosReq({
    url: '/userLogin/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 更新用户登录信息状态
export const updateUserLoginStatus = async (params: any) => {
  return axiosReq({
    url: '/userLogin/status',
    method: 'put',
    data: params,
  });
};
