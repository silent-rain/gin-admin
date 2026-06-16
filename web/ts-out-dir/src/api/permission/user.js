export const getUserInfo = async () => {
    return axiosReq({
        url: '/user/info',
        method: 'get',
        isParams: true,
        data: {},
    });
};
export const getAllUser = async () => {
    return axiosReq({
        url: '/user/all',
        method: 'get',
        isParams: true,
        data: {},
    });
};
export const getUserList = async (params) => {
    return axiosReq({
        url: '/user/list',
        method: 'get',
        isParams: true,
        data: params,
    });
};
export const addUser = async (params) => {
    return axiosReq({
        url: '/user/add',
        method: 'post',
        data: params,
    });
};
export const updateUser = async (params) => {
    return axiosReq({
        url: '/user/update',
        method: 'put',
        data: params,
    });
};
export const deleteUser = async (params) => {
    return axiosReq({
        url: '/user/delete',
        method: 'delete',
        data: params,
    });
};
export const batchDeleteUser = async (params) => {
    return axiosReq({
        url: '/user/batchDelete',
        method: 'delete',
        data: params,
    });
};
export const updateUserStatus = async (params) => {
    return axiosReq({
        url: '/user/updateStatus',
        method: 'put',
        data: params,
    });
};
export const updateUserPwd = async (params) => {
    return axiosReq({
        url: '/user/updatePwd',
        method: 'put',
        data: params,
    });
};
export const resetUserPwd = async (params) => {
    return axiosReq({
        url: '/user/resetPwd',
        method: 'put',
        data: params,
    });
};
export const updatePhone = async (params) => {
    return axiosReq({
        url: '/user/updatePhone',
        method: 'put',
        data: params,
    });
};
export const updateEmail = async (params) => {
    return axiosReq({
        url: '/user/updateEmail',
        method: 'put',
        data: params,
    });
};
