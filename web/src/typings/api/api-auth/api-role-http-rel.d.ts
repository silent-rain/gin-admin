/** 角色与Http协议接口关联 */
export interface ApiRoleHttpRel {
  id: number;
  role_id: number;
  api_id: number;
  created_at: string;
  updated_at: string;
}

export interface ApiRoleHttpRelRsp {
  data_list: ApiRoleHttpRel[];
  tatol: number;
}
