package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建默认路由引擎
	r := gin.Default()

	//处理GET请求，路径为/hello
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	//运行引擎，默认端口为8080
	r.Run(":9090")
}
