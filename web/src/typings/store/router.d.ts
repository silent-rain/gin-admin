/*
 * 声明.d.ts文件规范
 * 导出的类型以大写开头
 * 对象：config
 * 数组：options
 * 枚举：emu
 * 函数：Fn
 * 属性：props
 * 实例：instance
 * */

/* router */
import type { RouteRecordRaw } from 'vue-router';

export interface rawConfig {
  hidden?: boolean;
  alwaysShow?: boolean;
  code?: number;
  name?: string;
  fullPath?: string;
  path?: string;
  meta?: {
    title: string;
    icon?: string;
    affix?: boolean;
    activeMenu?: string;
    breadcrumb?: boolean;
    roles?: Array<number>;
    elSvgIcon?: string;
    code?: number;
    cachePage?: boolean;
    leaveRmCachePage?: boolean;
    closeTabRmCache?: boolean;
  };
  children?: RouterOptions;
  redirect?: string;
}
export type RouteRawConfig = RouteRecordRaw & rawConfig;
export type RouterTypes = Array<rawProp>;


export {};
