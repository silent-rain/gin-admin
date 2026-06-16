import type { User } from '@/typings/api/permission/user'
import type { Role } from '~/api/system/role'

export interface UserTry {
  token: string
  getUserInfo: boolean
  userInfo: User
  userId?: number
  userAvatar?: string
  roles: Role[]
  codes: number[]
}
