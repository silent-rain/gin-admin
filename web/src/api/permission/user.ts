/* 用户
 */

// 获取用户信息
export async function getUserInfo() {
  return axiosReq({
    url: '/user/info',
    method: 'get',
    isParams: true,
    data: {},
  })
}

// 获取所有用户列表
export async function getAllUser() {
  return axiosReq({
    url: '/user/all',
    method: 'get',
    isParams: true,
    data: {},
  })
}

// 获取用户列表
export async function getUserList(params: any) {
  return axiosReq({
    url: '/user/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加用户
export async function addUser(params: any) {
  return axiosReq({
    url: '/user/add',
    method: 'post',
    data: params,
  })
}

// 更新用户详情信息
export async function updateUser(params: any) {
  return axiosReq({
    url: '/user/update',
    method: 'put',
    data: params,
  })
}

// 删除用户
export async function deleteUser(params: any) {
  return axiosReq({
    url: '/user/delete',
    method: 'delete',
    data: params,
  })
}
// 批量删除用户
export async function batchDeleteUser(params: any) {
  return axiosReq({
    url: '/user/batchDelete',
    method: 'delete',
    data: params,
  })
}

// 更新用户状态
export async function updateUserStatus(params: any) {
  return axiosReq({
    url: '/user/updateStatus',
    method: 'put',
    data: params,
  })
}
// 更新用户密码
export async function updateUserPwd(params: any) {
  return axiosReq({
    url: '/user/updatePwd',
    method: 'put',
    data: params,
  })
}
// 重置用户密码
export async function resetUserPwd(params: any) {
  return axiosReq({
    url: '/user/resetPwd',
    method: 'put',
    data: params,
  })
}
// 更新用户手机号码
export async function updatePhone(params: any) {
  return axiosReq({
    url: '/user/updatePhone',
    method: 'put',
    data: params,
  })
}
// 更新用户邮箱
export async function updateEmail(params: any) {
  return axiosReq({
    url: '/user/updateEmail',
    method: 'put',
    data: params,
  })
}
