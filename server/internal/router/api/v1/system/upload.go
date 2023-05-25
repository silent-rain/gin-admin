/*上传管理*/
package system

import (
	"github.com/silent-rain/gin-admin/internal/app/system/controller"

	"github.com/gin-gonic/gin"
)

// InitUploadRouter 初始化上传管理路由
func InitUploadRouter(group *gin.RouterGroup) {
	// 文件上传
	upload := group.Group("/upload")
	{
		// 上传用户头像
		upload.POST("/avatar", controller.NewUploadController().Avatar)
		// 上传图片
		upload.POST("/image", controller.NewUploadController().Image)
		// 上传图片列表
		upload.POST("/images", controller.NewUploadController().Images)
	}
}
