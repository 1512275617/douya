package router

import (
	"douya/admin/handler"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {
	adminGroup := g.Group("/admin")

	// 健康检查
	adminGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// hello world
	adminGroup.GET("/say", handler.HelloHandler{}.GetHelloHandler)
}
