package router

import (
	"task4/common"
	"task4/controller"

	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello blog",
		})
	})
	v1 := r.Group("/api")
	{
		v1.GET("/articles", controller.GetArticles)
		v1.GET("/article/:id", controller.GetArticle)
		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)

		// 需要认证的路由组
		auth := v1.Group("", common.AuthMiddleware())
		{
			// 创建文章接口
			auth.POST("/article", controller.CreateArticle)
			// 更新文章接口
			auth.PUT("/article/:id", controller.UpdateArticle)
			// 删除文章接口
			auth.DELETE("/article/:id", controller.DeleteArticle)
		}
	}
}
