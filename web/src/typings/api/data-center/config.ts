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

// 网站配置
export interface WebsiteConfig {
  website_title: Config;
  website_intro: Config;
  website_logo: Config;
  website_seo_title: Config;
  website_seo_desc: Config;
  website_copyright: Config;
  website_keyword: Config;
  website_desc: Config;
  website_company_address: Config;
  website_phone: Config;
  website_email: Config;
  website_qq: Config;
  website_filing_number: Config;
  website_propaganda: Config;
  website_tags: Config;
}
