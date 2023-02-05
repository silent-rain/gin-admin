/**角色菜单关系 */
export interface RoleMenuRel {
  id: number;
  role_id: number;
  menu_id: number;
  created_at: string;
  updated_at: string;
}

export interface RoleMenuRelListRsp {
  data_list: RoleMenuRel[];
  tatol: number;
}
