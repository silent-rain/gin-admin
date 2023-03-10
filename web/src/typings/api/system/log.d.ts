/* 日志
 */
export interface HttpLog {
  id: number;
  user_id: number;
  trace_id: string;
  span_id: string;
  status_code: number;
  method: string;
  path: string;
  query: string;
  body: string;
  remote_addr: string;
  user_agent: number;
  cost: string;
  htpp_type: string;
  note: string;
  created_at: string;
}

export interface HttpLogListRsp {
  code: number;
  msg: string;
  data_list: HttpLog[];
  tatol: number;
}

export interface SystemLog {
  id: number;
  user_id: number;
  trace_id: string;
  level: string;
  caller_line: string;
  error_code: number;
  error_msg: string;
  msg: string;
  extend: string;
  note: string;
  created_at: string;
}

export interface SystemLogListRsp {
  code: number;
  msg: string;
  data_list: SystemLog[];
  tatol: number;
}

export interface WebLog {
  id: number;
  user_id: number;
  nickname: string;
  trace_id: string;
  os_type: number;
  error_type: number;
  level: string;
  caller_line: string;
  url: string;
  msg: string;
  stack: string;
  note: string;
  created_at: string;
}

export interface WebLogListRsp {
  code: number;
  msg: string;
  data_list: WebLog[];
  tatol: number;
}
