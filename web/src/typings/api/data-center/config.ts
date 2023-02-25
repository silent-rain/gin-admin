/* 配置
 */

export interface Config {
  id: number;
  parent_id: number;
  name: string;
  key: string;
  value: string;
  sort: number;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
  children: Config[];
}
export interface ConfigListRsp {
  data_list: Config[];
  tatol: number;
}
