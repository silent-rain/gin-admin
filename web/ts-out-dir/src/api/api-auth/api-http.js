import axiosReq from '@/utils/axios-req';
export const getAllApiHttpTree = async () => {
    return axiosReq({
        url: '/apiHttp/allTree',
        method: 'get',
        isParams: true,
        data: {},
    });
};
export const getApiHttpTree = async (params) => {
    return axiosReq({
        url: '/apiHttp/tree',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addApiHttp = async (params) => {
    return axiosReq({
        url: '/apiHttp/add',
        method: 'post',
        data: params,
    });
};
export const updateApiHttp = async (params) => {
    return axiosReq({
        url: '/apiHttp/update',
        method: 'put',
        data: params,
    });
};
export const deleteApiHttp = async (params) => {
    return axiosReq({
        url: '/apiHttp/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteApiHttp = async (params) => {
    return axiosReq({
        url: '/apiHttp/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateApiHttpStatus = async (params) => {
    return axiosReq({
        url: '/apiHttp/updateStatus',
        method: 'put',
        data: params,
    });
};
