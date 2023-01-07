/*
 * @Author: silent-rain
 * @Date: 2023-01-07 18:44:38
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 19:01:20
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/handler/handler.go
 * @Descripttion: 服务逻辑
 */
package handler

import (
	"net/http"

	"gin-admin/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// Ping 服务健康检查
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{})
}

// SayHello 接口测试
func SayHello(c *gin.Context) {
	name := c.Param("name")
	response.New(c).WithData("hello," + name).Json()
}
