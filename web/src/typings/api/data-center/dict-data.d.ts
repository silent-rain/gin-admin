/** 字典数据管理
 *
 */

export interface DictData {
  id: number;
  dict_id: number;
  name: string;
  value: string;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
}
export interface DictDataListRsp {
  code: number;
  msg: string;
  data_list: DictData[];
  tatol: number;
}
