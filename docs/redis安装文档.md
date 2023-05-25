# Redis 安装与配置

## 安装
```
sudo apt install redis
```


## 配置
```
sudo vim /etc/redis/redis.conf


# 配置远程连接
# 修改bind 127.0.0.1
bind 0.0.0.0

# 修改端口
port 6379

# 配置密码
# 取消注释 requirepass foobared, foobared 即密码，可自行修改
requirepass foobared
```

## 检查Redis服务器系统进程
```
ps -aux|grep redis
```

## 查看Redis端口状态
```
netstat -nlt|grep 6379
```

## 开机自启
```
sudo systemctl start redis-server.service
```

## 连接 redis 服务端
```
# 连接
redis-cli

# 认证
auth
```