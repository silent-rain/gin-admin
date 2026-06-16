import axiosReq from '@/utils/axios-req';
export const getCaptcha = async () => {
    return axiosReq({
        url: '/captcha',
        method: 'get',
        bfLoading: false,
        isAlertErrorMsg: false,
        isParams: true,
        data: {},
    });
};
export const captchaVerify = async (params) => {
    return axiosReq({
        url: '/captcha/verify',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const register = async (params) => {
    return axiosReq({
        url: '/register',
        method: 'post',
        data: params,
    });
};
export const login = async (data) => {
    return axiosReq({
        url: '/login',
        data,
        method: 'post',
    });
};
export const logout = async () => {
    return axiosReq({
        url: '/logout',
        method: 'post',
    });
};
