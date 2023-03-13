package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建默认引擎
	r := gin.Default()
	//请求编写
	//查询
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})
	//创建
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	//修改
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	//删除
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	//启动HTTP服务
	r.Run()
}
