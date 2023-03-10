/* 用户
 */

export interface User {
  id: number;
  realname: string;
  nickname: string;
  gender: number;
  age: number;
  birthday: string;
  avatar: string;
  phone: string;
  email: string;
  intro: string;
  note: string;
  password: string;
  password2: string;
  sort: number;
  status: number;
  created_at: string;
  updated_at: string;
  roles: any[];
  role_ids: number[];
  captcha_id: string;
  captcha: string;
}

export interface UserListRsp {
  code: number;
  msg: string;
  data_list: User[];
  tatol: number;
}
