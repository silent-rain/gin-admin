/*
 * @Author: silent-rain
 * @Date: 2023-01-08 15:57:53
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-08 16:03:08
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/web/src/utils/md5.ts
 * @Descripttion: MD5 加密
 */
import { secret } from "@/utils/constant";
import md5 from "js-md5";

// MD5 加密
export const md5Encode = (v: string) => {
  return md5(secret + v);
};
