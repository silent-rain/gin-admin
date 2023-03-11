/* Http协议接口管理
 */

export interface ApiHttp {
  id: number;
  parent_id: number | undefined;
  name: string;
  method: string;
  uri: string;
  note: string;
  status: number;
  created_at: string;
  updated_at: string;
  children: ApiHttp[];
}

export interface ApiHttpTreeRsp {
  code: number;
  msg: string;
  data_list: ApiHttp[];
  tatol: number;
}
