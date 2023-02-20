/*浏览器*/
package plugin

import (
	"fmt"
	"os/exec"
	"runtime"

	"gin-admin/internal/pkg/conf"

	"go.uber.org/zap"
)

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

// Open 打开浏览器
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

// RegisterOpenBrowser 服务启动后在浏览器中打开 URI
func RegisterOpenBrowser() {
	if !conf.Instance().Server.Plugin.EnableOpenBrowser {
		return
	}
	if err := Open(conf.Instance().Server.Plugin.OpenBrowserUrl); err != nil {
		zap.S().Error(err)
		return
	}
}
