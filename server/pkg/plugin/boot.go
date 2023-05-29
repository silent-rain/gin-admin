// Package plugin 服务启动
package plugin

import (
	"fmt"

	"github.com/silent-rain/gin-admin/internal/global"
	"github.com/silent-rain/gin-admin/pkg/color"
)

// see https://patorjk.com/software/taag/#p=testall&f=Graffiti&t=go-gin-api
const logo = `
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝
`

// RegisterLogo 服务启动后显示 logo
func RegisterLogo() {
	if !global.Instance().Config().Server.Plugin.EnableLogo {
		return
	}
	fmt.Println(color.Blue(logo))
}

// RegisterAddr 服务启动后显示 IP 地址
func RegisterAddr() {
	cfg := global.Instance().Config().Server
	if !cfg.Plugin.EnableLogo {
		return
	}
	addr := cfg.ServerAddress()
	fmt.Println(color.Blue(fmt.Sprintf("➜  Local:   http://%s", addr)))
}
