/* 用户API接口Token令牌表
 */

export interface UserApiToken {
  id: number;
  user_id: number;
  nickname: string;
  permission: string;
  token: string;
  passphrase: string;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
}

export interface UserApiTokenListRsp {
  data_list: UserApiToken[];
  tatol: number;
}
