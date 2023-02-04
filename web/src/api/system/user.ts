/*用户
 */

// 获取用户信息
export const getUserInfo = async (params: any) => {
  return axiosReq({
    url: '/user/info',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 获取所有用户列表
export const getAllUser = async () => {
  return axiosReq({
    url: '/user/all',
    method: 'get',
    isParams: true,
    data: {},
  });
};

// 获取用户列表
export const getUserList = async (params: any) => {
  return axiosReq({
    url: '/user/list',
    method: 'get',
    isParams: true,
    data: params,
  });
};

// 添加用户
export const addUser = async (params: any) => {
  return axiosReq({
    url: '/user/add',
    method: 'post',
    data: params,
  });
};

// 更新用户详情信息
export const updateUser = async (params: any) => {
  return axiosReq({
    url: '/user/update',
    method: 'put',
    data: params,
  });
};

// 删除用户
export const deleteUser = async (params: any) => {
  return axiosReq({
    url: '/user/delete',
    method: 'delete',
    data: params,
  });
};
// 批量删除用户
export const batchDeleteUser = async (params: any) => {
  return axiosReq({
    url: '/user/batchDelete',
    method: 'delete',
    data: params,
  });
};

// 更新用户状态
export const updateUserStatus = async (params: any) => {
  return axiosReq({
    url: '/user/status',
    method: 'put',
    data: params,
  });
};
// 更新用户密码
export const updateUserPwd = async (params: any) => {
  return axiosReq({
    url: '/user/updatePwd',
    method: 'put',
    data: params,
  });
};
// 重置用户密码
export const resetUserPwd = async (params: any) => {
  return axiosReq({
    url: '/user/resetPwd',
    method: 'put',
    data: params,
  });
};
// 更新用户手机号码
export const updatePhone = async (params: any) => {
  return axiosReq({
    url: '/user/updatePhone',
    method: 'put',
    data: params,
  });
};
// 更新用户邮箱
export const updateEmail = async (params: any) => {
  return axiosReq({
    url: '/user/updateEmail',
    method: 'put',
    data: params,
  });
};
