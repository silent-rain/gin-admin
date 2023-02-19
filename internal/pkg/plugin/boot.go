/*服务启动*/
package plugin

import (
	"fmt"

	"gin-admin/internal/pkg/conf"
	"gin-admin/pkg/color"
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

// ShowLogo 服务启动后显示 logo
func ShowLogo() {
	if !conf.Instance().Server.Base.EnableLogo {
		return
	}
	fmt.Println(color.Blue(logo))
}
