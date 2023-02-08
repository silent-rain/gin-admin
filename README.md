# Gin-Admin 后台管理系统

## 开发环境
- 系统版本: deepin v20.8
- go 版本: v1.18
- node 版本: v16.19.0
- yarn 版本: v1.22.18
- 数据库 MariaDB 版本: v10.3.36

## 后端框架功能列表
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
- [x] API 文档
  - [x] ApiPost 接口工具
  - [x] 内置接口文档
- [ ] 定时任务调度
  - [ ] 即时任务
  - [ ] 定时任务

## 系统功能列表
- [ ] 系统管理
  - [x] 注册/登录/退出/验证码
  - [x] 用户管理
    - [ ] 用户中心
    - [x] 头像上传
  - [x] 角色管理
  - [x] 菜单管理
- [ ] 前端动态路由/权限
- [ ] 系统设置
  - [ ] 系统配置管理


## 待处理
- 便捷工具，列展示功能待优化
  - 窗口跳动问题 - 高度滚动条导致的
- 用户页面
  - 导入

## 项目编译&运行
### 后端
此后端内嵌web静态资源，打包后可直接访问。
- 热重启, 修改代码后自动编译运行

```shell
# 项目根目录
air
```
- 项目调式运行

```shell
cd cmd
go run main.go
```
- 项目编译&运行

```shell
cd cmd
# 编译
go build -o ./main .
# 添加执行权限
chmod 755 main
# 运行
./main
# 后端访问地址
http://127.0.0.1:8080/ping
```
### 前端编译&运行
- 安装依赖

```shell
cd web
yarn install
```
- 调试运行

```shell
cd web
yarn run dev
# 访问前端
http://localhost:5005/
```
- 编译

```shell
cd web
yarn build
```

## API 接口文档
本地API文档，需要联网加载资源
[本地API文档](http://127.0.0.1:8080/docs/api/v1/index.html)

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
- [参考项目](http://manage.pro.layui.javaweb.vip/index)

