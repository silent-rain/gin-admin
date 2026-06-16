/* 注册/登录/登出/验证码
 */
import axiosReq from '@/utils/axios-req'

// 获取验证码
export async function getCaptcha() {
  return axiosReq({
    url: '/captcha',
    method: 'get',
    bfLoading: false,
    isAlertErrorMsg: false,
    isParams: true,
    data: {},
  })
}

// 验证码验证
export async function captchaVerify(params: any) {
  return axiosReq({
    url: '/captcha/verify',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 注册
export async function register(params: any) {
  return axiosReq({
    url: '/register',
    method: 'post',
    data: params,
  })
}

// 登录
export async function login(data: any) {
  return axiosReq({
    url: '/login',
    data,
    method: 'post',
  })
}

// 退出登录
export async function logout() {
  return axiosReq({
    url: '/logout',
    method: 'post',
  })
}
