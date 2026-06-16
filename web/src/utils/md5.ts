import { md5 } from 'js-md5'
import { SECRET } from '@/constant/permission/auth'

// MD5 加密
export function md5Encode(v: string) {
  return md5(SECRET + v)
}
