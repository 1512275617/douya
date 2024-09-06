package app

import (
	"douya/app/handler"
	"douya/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	appGroup := g.Group("/douya")

	// 健康检查1
	appGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// 获取token
	appGroup.GET("/token", handler.NewTokenHandler().GenerateToken)

	// 鉴权例子
	appGroup.GET("/haha", utils.Handler(true), handler.NewTokenHandler().Haha)
}
