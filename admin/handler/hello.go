package handler

import (
	"douya/admin/service"
	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
}

func (h HelloHandler) GetHelloHandler(ctx *gin.Context) {
	info, err := service.Hello(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code": 400,
			"msg":  "请求异常",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": info,
	})
}
