# 问答

## Failed to find the "go" binary from PATH
- 安装 go 插件后报错

```text
Every time I restart vscode, it shows: Failed to find the "go" binary in either。GOROOT() or PATH(/usr/bin:/bin:/usr/sbin:/sbin). Check PATH, or Install Go and reload the window.

It seems the PATH(/usr/bin:/bin:/usr/sbin:/sbin) from the message is not my actual path:
```
- 解决方案：

在 settings.json 配置文件中指定路径
```
{
    "go.goroot": "/opt/devSoft/go"
}
```