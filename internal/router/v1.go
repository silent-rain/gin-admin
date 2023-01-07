/*
 * @Author: silent-rain
 * @Date: 2023-01-06 00:26:00
 * @LastEditors: silent-rain
 * @LastEditTime: 2023-01-07 19:00:16
 * @company:
 * @Mailbox: silent_rains@163.com
 * @FilePath: /gin-admin/internal/router/v1.go
 * @Descripttion:
 */
/**Api version 1 路由
 */
package router

import (
	"gin-admin/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewApiV1(engine *gin.Engine) {
	v1 := engine.Group("api/v1")
	// 接口测试
	v1.GET("/sayHello/:name", handler.SayHello)

	// 服务1
	server1 := v1.Group("server1")
	{
		server1.GET("/sayHello/:name", SayHello)
		server1.GET("/test/:id/:name", getUser)
	}

}

//http://localhost:9090/test/1/dong
func getUser(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	json := gin.H{
		"data": id,
		"name": name,
	}
	c.JSON(http.StatusOK, json)
}

//http://localhost:9090/sayHello/dong
func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello,"+name)
	//c.String(http.StatusOK, "change,"+name)
}
