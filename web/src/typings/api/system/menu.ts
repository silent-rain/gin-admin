/*菜单
 */

export interface Menu {
  id: number;
  parent_id: number;
  title: string;
  icon: string;
  menu_type: number;
  open_type: number;
  path: string;
  component: string;
  link: string;
  target: string;
  permission: string;
  hide: number;
  sort: number;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
  children: Menu[];
  _rawData: Menu;
}
export interface MenuListRsp {
  data_list: Menu[];
  tatol: number;
}
