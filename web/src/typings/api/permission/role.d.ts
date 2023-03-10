/* 角色
 */
export interface Role {
  id: number;
  name: string;
  sort: number;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
}

export interface RoleListRsp {
  code: number;
  msg: string;
  data_list: Role[];
  tatol: number;
}
