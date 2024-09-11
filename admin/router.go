package admin

import (
	"douya/admin/handler"
	"douya/utils/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	g.Use(cors.Cors())
	adminGroup := g.Group("/admin")

	// 健康检查
	adminGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// hello world
	adminGroup.GET("/say", handler.NewHelloHandler().GetHelloHandler)
}
