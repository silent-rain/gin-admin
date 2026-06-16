import axiosReq from '@/utils/axios-req';
export const getAllConfigTree = async () => {
    return axiosReq({
        url: '/config/allTree',
        method: 'get',
        isParams: true,
        data: {},
    });
};
export const getConfigList = async (params) => {
    return axiosReq({
        url: '/config/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const getConfigChildrensByKey = async (params) => {
    return axiosReq({
        url: '/config/childrensByKey',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const getConfigTree = async (params) => {
    return axiosReq({
        url: '/config/tree',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const getConfigInfo = async (params) => {
    return axiosReq({
        url: '/config/info',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addConfig = async (params) => {
    return axiosReq({
        url: '/config/add',
        method: 'post',
        data: params,
    });
};
export const updateConfig = async (params) => {
    return axiosReq({
        url: '/config/update',
        method: 'put',
        data: params,
    });
};
export const batchUpdateConfig = async (params) => {
    return axiosReq({
        url: '/config/batchUpdate',
        method: 'put',
        data: params,
    });
};
export const deleteConfig = async (params) => {
    return axiosReq({
        url: '/config/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteConfig = async (params) => {
    return axiosReq({
        url: '/config/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateConfigStatus = async (params) => {
    return axiosReq({
        url: '/config/updateStatus',
        method: 'put',
        data: params,
    });
};
export const getWebSiteConfigList = async () => {
    return axiosReq({
        url: '/config/webSiteConfigList',
        method: 'get',
        isParams: true,
        data: {},
    });
};
