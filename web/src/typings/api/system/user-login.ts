/**用户登录管理
 *
 */

export interface UserLogin {
  id: number;
  user_id: number;
  nickname: string;
  remote_addr: string;
  user_agent: string;
  status: number;
  created_at: string;
  updated_at: string;
}

export interface UserLoginListRsp {
  data_list: UserLogin[];
  tatol: number;
}
