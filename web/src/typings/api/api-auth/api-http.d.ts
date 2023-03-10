/* Http协议接口管理
 */

export interface ApiHttp {
  id: number;
  name: string;
  method: string;
  uri: string;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
}

export interface ApiHttpListRsp {
  code: number;
  msg: string;
  data_list: ApiHttp[];
  tatol: number;
}
