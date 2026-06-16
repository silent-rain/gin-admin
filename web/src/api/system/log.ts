/**
 * 日志
 *
 */
import { WebErrorType } from '@/constant/system/log'
import { useBasicStore } from '@/store/basic'
import axiosReq from '@/utils/axios-req'

/**
 * 后端日志
 *
 */
// 获取网络请求日志列表
export async function getHttpLogList(params: any) {
  return axiosReq({
    url: '/httpLog/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 获取网络请求日志 body 信息
export async function getHttpLogBody(params: any) {
  return axiosReq({
    url: '/httpLog/body',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 获取系统日志列表
export async function getSystemLogList(params: any) {
  return axiosReq({
    url: '/systemLog/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

/**
 * WEB日志
 *
 */
// 获取 WEB 日志列表
export async function getWebList(params: any) {
  return axiosReq({
    url: '/webLog/list',
    method: 'get',
    isParams: true,
    data: params,
  })
}

// 添加 WEB 代码日志
export async function addWebCodeLog(params: any) {
  if (!params.level) {
    params.level = 'ERROR'
  }
  params.error_type = WebErrorType.Code
  params.os_type = useBasicStore().osType
  params.url = window.location.href
  return axiosReq({
    url: '/webLog/add',
    method: 'post',
    data: params,
  })
}

// 添加 WEB API 日志
export async function addWebApiLog(params: any) {
  if (!params.level) {
    params.level = 'ERROR'
  }
  params.error_type = WebErrorType.Api
  params.os_type = useBasicStore().osType
  params.url = window.location.href
  return axiosReq({
    url: '/webLog/add',
    method: 'post',
    data: params,
  })
}
