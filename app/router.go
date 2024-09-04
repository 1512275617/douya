package app

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	appGroup := g.Group("/douya")

	// 健康检查
	appGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
