import axiosReq from '@/utils/axios-req';
export const getDictDataList = async (params) => {
    return axiosReq({
        url: '/dictData/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addDictData = async (params) => {
    return axiosReq({
        url: '/dictData/add',
        method: 'post',
        data: params,
    });
};
export const updateDictData = async (params) => {
    return axiosReq({
        url: '/dictData/update',
        method: 'put',
        data: params,
    });
};
export const deleteDictData = async (params) => {
    return axiosReq({
        url: '/dictData/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteDictData = async (params) => {
    return axiosReq({
        url: '/dictData/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateDictDataStatus = async (params) => {
    return axiosReq({
        url: '/dictData/updateStatus',
        method: 'put',
        data: params,
    });
};
