/** 字典维度管理
 *
 */

export interface Dict {
  id: number;
  name: string;
  code: string;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
}
export interface DictListRsp {
  code: number;
  msg: string;
  data_list: Dict[];
  tatol: number;
}
