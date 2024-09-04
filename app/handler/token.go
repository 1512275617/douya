package handler

import (
	"douya/utils"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

func (handler *TokenHandler) GenerateToken(ctx *gin.Context) {
	token, err := utils.GenerateToken()
	if err != nil {
		ctx.JSON(400, gin.H{
			"code": 400,
			"msg":  "请求异常",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": token,
	})
}

func (handler *TokenHandler) Haha(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": "恭喜校验token成功了!",
	})
}
