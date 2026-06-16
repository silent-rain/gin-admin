import axiosReq from '@/utils/axios-req';
export const getUserApiTokenList = async (params) => {
    return axiosReq({
        url: '/userApiToken/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addUserApiToken = async (params) => {
    return axiosReq({
        url: '/userApiToken/add',
        method: 'post',
        data: params,
    });
};
export const updateUserApiToken = async (params) => {
    return axiosReq({
        url: '/userApiToken/update',
        method: 'put',
        data: params,
    });
};
export const deleteUserApiToken = async (params) => {
    return axiosReq({
        url: '/userApiToken/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteUserApiToken = async (params) => {
    return axiosReq({
        url: '/userApiToken/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateUserApiTokenStatus = async (params) => {
    return axiosReq({
        url: '/userApiToken/updateStatus',
        method: 'put',
        data: params,
    });
};
