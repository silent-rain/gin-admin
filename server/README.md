# Gin-Admin 后台管理系统
- `Gin-Admin`后端是基于[Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的框架功能，使用简单，致力于快速的业务研发。
- 前端是基于[Vue框架模板](https://github.com/jzfai/vue3-admin-plus)进行设计的前端管理框架，封装了常用的权限管理、系统管理功能，致力于一套开箱即用的权限系统。
- 同时支持将web服务内置于go的编译包中，编译后可直接访问前后端。

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
  - [x] mysql 全局实例
    - [x] 只读实例
    - [x] 读写实例
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
  - [ ] 登录表 鉴权
  - [x] API 鉴权
  - [ ] API Token 鉴权
  - [x] [rate](https://golang.org/x/time/rate) 接口限流 
- [x] API 文档
  - [x] ApiPost 接口工具
  - [x] 内置接口文档
- [ ] 定时任务调度
  - [ ] 即时任务
  - [ ] 定时任务
- [x] 优雅关机
  - [x] HTTP 服务
  - [x] Mysql 数据库
- [ ] 插件
  - [x] 服务启动logo
  - [x] [pprof](https://github.com/gin-contrib/pprof) 性能剖析工具
  - [x] [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
  - [ ] [Swagger](https://github.com/swaggo/gin-swagger) 接口文档
  - [x] 服务启动后打开浏览器
## 系统功能列表
- [ ] 权限管理
  - [x] 注册/登录/退出/验证码
  - [x] 用户管理
    - [ ] 用户中心
    - [x] 头像上传
  - [x] 角色管理
  - [x] 菜单管理
    - [ ]    el_svg_icon: string;
    - [ ]    icon: string;
- [x] 前端权限
  - [x] 动态路由
  - [x] 按钮权限权限
  - [ ] 后端权限配置提交时校验
- [ ] 系统管理
  - [x] 系统配置管理
  - [ ] 可视化系统配置管理
  - [ ] 请求日志
    - [ ] ui
  - [ ] 系统日志
    - [ ] ui
  - [ ] 登录管理

## 待处理
- 便捷工具，列展示功能待优化
  - 窗口跳动问题 - 高度滚动条导致的
- 用户页面
  - 导入
  - 角色权限优化

1. 支持 panic 异常时邮件通知
Excel
缓存
redis
订阅模式
2. 支持 trace 项目内部链路追踪 
3. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
4. 支持 [cron](https://github.com/jakecoffman/cron) 定时任务，在后台可界面配置
5. 支持 [websocket](https://github.com/gorilla/websocket) 实时通讯，在后台有界面演示
6. /app/Kit/ ※ 自定义系统服务，包含第三方服务和系统功能服务
/app/Ruler/Task/ ※ app全局定时任务，默认20s精度




## 项目编译&运行
### 后端
此后端内嵌web静态资源，打包后可直接访问。
#### 开发模式
- 热重启, 修改代码后自动编译运行

```shell
# 进入后端项目
cd server
# 执行air指令即可启动项目
air
```
- 手动调式运行

```shell
cd server/cmd
go run main.go
```
#### 发布模式
- 项目编译

```shell
cd server/cmd
# 编译
go build -o ./main .
```

- 运行
```shell
# 添加执行权限
chmod 755 main
# 运行
./main
```
#### 访问服务
- [接口连接测试](http://127.0.0.1:8080/api/ping)
- [内嵌前端](http://127.0.0.1:8080/)
- [pprof 性能剖析工具](http://127.0.0.1:8080/debug/pprof/)
- [Prometheus 监控指标](http://127.0.0.1:8080/metrics)

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
- [为 Go 应用添加 Prometheus 监控指标](https://blog.csdn.net/weixin_40046357/article/details/120620433)

### go-api 设计参考
https://github.com/flipped-aurora/gin-vue-admin/tree/main/server
https://github.com/go-admin-team/go-admin
https://github.com/GoAdminGroup/go-admin
https://github.com/eddycjy/go-gin-example
https://github.com/xinliangnote/go-gin-api
https://github.com/bg-vc/vc-gin-api
https://github.com/fyonecon/ginlaravel
https://github.com/mritd/ginmvc
### 日志
https://zhuanlan.zhihu.com/p/430224518
https://zhuanlan.zhihu.com/p/553995164

### 定时任务
https://www.cnblogs.com/hsyw/p/16006799.html
https://blog.csdn.net/JineD/article/details/121214032
https://zhuanlan.zhihu.com/p/402210167