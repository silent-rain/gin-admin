/*注册/登录/登出/验证码
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

// 注册
export const register = async (params: any) => {
  return axiosReq({
    url: '/register',
    method: 'post',
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
