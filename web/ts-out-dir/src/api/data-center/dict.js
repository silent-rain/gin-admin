import axiosReq from '@/utils/axios-req';
export const getDictList = async (params) => {
    return axiosReq({
        url: '/dict/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addDict = async (params) => {
    return axiosReq({
        url: '/dict/add',
        method: 'post',
        data: params,
    });
};
export const updateDict = async (params) => {
    return axiosReq({
        url: '/dict/update',
        method: 'put',
        data: params,
    });
};
export const deleteDict = async (params) => {
    return axiosReq({
        url: '/dict/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteDict = async (params) => {
    return axiosReq({
        url: '/dict/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateDictStatus = async (params) => {
    return axiosReq({
        url: '/dict/updateStatus',
        method: 'put',
        data: params,
    });
};
