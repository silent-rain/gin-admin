import { User } from '@/typings/api/permission/user';
import { Role } from '~/api/system/role';
import { Menu } from '~/api/system/menu';

export interface UserTry {
  token: string;
  getUserInfo: boolean;
  userInfo: User;
  userId?: number;
  userAvatar?: string;
  roles: Role[];
  codes: number[];
}
