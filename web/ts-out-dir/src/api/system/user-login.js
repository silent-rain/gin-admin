import axiosReq from '@/utils/axios-req';
export const getUserLoginList = async (params) => {
    return axiosReq({
        url: '/userLogin/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const updateUserLoginStatus = async (params) => {
    return axiosReq({
        url: '/userLogin/updateStatus',
        method: 'put',
        data: params,
    });
};
