/*用户
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
  sort: number;
  status: number;
  created_at: string;
  updated_at: string;
  role_ids: number[];
  _rawData: User;
}

export interface UserListRsp {
  data_list: User[];
  tatol: number;
}
