<!--
 * @Author: silent-rain
 * @Date: 2023-01-05 00:20:01
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-14 17:18:48
 * @company: 
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/README.md
 * @Descripttion: 后台管理系统
-->
# Gin-Admin 后台管理系统

## 开发环境
- 系统版本: deepin v20.8
- go 版本: v1.18
- node 版本: v16.19.0
- yarn 版本: v1.22.18
- 数据库 MariaDB 版本: v10.3.36

## 功能列表
- [x] 热重启
- [x] 内嵌文件
- [x] 内嵌 web 服务
  - [x] vue3+vite+ts+elementUI 项目
- [x] API接口状态码
- [x] 自定义API接口返回结构
- [x] 数据库 gorm
  - [x] 数据库类型 sqlite3/mysql
  - [x] DB 全局实例
- [x] 日志 zap
  - [x] 本地滚动日志 - 文本日志
  - [x] 数据库日志 
    - [x] 结构化日志
    - [x] 单独封装入库
  - [x] trace 日志
  - [x] Gin 框架集成 zap日志库
- [ ] 中间件
  - [x] 跨域
  - [x] API 网络请求/响应日志
  - [x] JWT 令牌
  - [x] API 鉴权
  - [ ] API Token 鉴权
- [ ] API 文档
  - [ ] ApiPost 接口工具
  - [ ] 内置在线接口文档
- [ ] 定时任务调度
  - [ ] 即时任务
  - [ ] 定时任务
- [ ] 系统管理
  - [x] 注册/登录/退出/验证码
  - [x] 用户管理
  - [x] 角色管理
  - [ ] 菜单管理
    - [ ] 角色菜单管理


## vscode 插件
### koroFileHeader
- 文档注释插件
- 插件配置
```
// 自动添加文件头部注释
"fileheader.configObj": {
  "autoAdd": true, // 自动添加头部注释开启才能自动添加`
  "autoAlready": true, // 默认开启`
},
// 文件头部注释
"fileheader.customMade": {
  "version": "V1.0.0", //版本号
  "Author": "lilun", //作者
  "Date": "Do not edit", //文件创建创建时间
  "LastEditors": "lilun", //最后编辑作者
  "LastEditTime": "Do not Edit", //最后编辑时间
  "company": "轩田科技", //公司名称
  "Mailbox": "lilun@sharetek.com.cn",//邮箱
  "FilePath": "Do not edit", // 文件在项目中的相对路径 自动更新
  "Descripttion": "", //文本描述
}, 
//函数注释
"fileheader.cursorMode": {
  "description": "", //方法描述
  "param ": "", //参数
  "return": ""
},
```
- 使用快捷键
  - 文件头部注释： 在当前编辑文件中使用快捷键即可生成文件头部注释。
    - ctrl+alt+i
  - 函数注释： 将光标放在函数行或者将光标放在函数上方的空白行 使用快捷键即可生成函数注释。
    - ctrl+alt+t

## 热重启
- install air
```
go install github.com/cosmtrek/air@latest
```
- init air
```
cd cmd
air init
```
- run
```
air
```


## 参考文档
- [前端框架模板](https://github.com/jzfai/vue3-admin-plus)
- [Gin框架基础](https://blog.csdn.net/qq_40229166/article/details/118807361)
- [Gorm](https://gorm.io/zh_CN/)
- [参考项目](http://manage.gin.elevue.easygoadmin.vip/system/user)

