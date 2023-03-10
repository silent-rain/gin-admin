/* 菜单
 */

export interface Menu {
  id: number;
  parent_id: number;
  title: string;
  icon: string;
  el_svg_icon: string;
  menu_type: number;
  open_type: number;
  path: string;
  name: string;
  component: string;
  redirect: string;
  link: string;
  target: string;
  permission: string;
  hidden: number;
  always_show: number;
  sort: number;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
  children: Menu[];
}
export interface MenuListRsp {
  code: number;
  msg: string;
  data_list: Menu[];
  tatol: number;
}

export interface ButtonPermission {
  permission: string; // 按钮权限标识
  disabled: number; // 是否禁用
}
