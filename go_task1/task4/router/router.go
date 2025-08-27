package router

import "github.com/gin-gonic/gin"

func SetRouters(r *gin.Engine) {
	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world ver",
		})
	})
}
