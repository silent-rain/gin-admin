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
  data_list: Role[];
  tatol: number;
}