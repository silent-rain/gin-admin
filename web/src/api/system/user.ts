/*
 * @Author: silent-rain
 * @Date: 2023-01-06 23:20:53
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-15 01:49:04
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/api/user.ts
 * @Descripttion: 用户
 */
import axiosReq from '@/utils/axios-req';

// 获取验证码
export const captcha = async (params: any) => {
  return axiosReq({
    url: '/captcha',
    method: 'get',
    bfLoading: false,
    isAlertErrorMsg: false,
    isParams: true,
    data: params,
  });
};

// 验证码验证
export const captchaVerify = async (params: any) => {
  return axiosReq({
    url: '/captcha/verify',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 登录
export const login = async (data: any) => {
  return axiosReq({
    url: '/login',
    data,
    method: 'post',
  });
};

// 退出登录
export const logout = async () => {
  return axiosReq({
    url: '/logout',
    method: 'post',
  });
};

// 获取用户信息
export const getUserInfo = async (params: any) => {
  return axiosReq({
    url: '/user/info',
    method: 'get',
    isParams: true,
    data: params,
  });
};
