/**
 * 用户登录管理
 *
 */

import axiosReq from '@/utils/axios-req'

// 获取用户登录信息列表
export async function getUserLoginList(params: any) {
  return axiosReq({
    url: '/userLogin/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 更新用户登录信息状态
export async function updateUserLoginStatus(params: any) {
  return axiosReq({
    url: '/userLogin/updateStatus',
    method: 'put',
    data: params,
  })
}
