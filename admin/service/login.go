package service

import (
	"douya/admin/entity"
	"douya/dao"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(c *gin.Context) {
	in := new(entity.LoginPassWordReq)
	err := c.ShouldBindJSON(&in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	//根据账号和密码查询用户
	sysUser, err := dao.GetUserNamePassword(in.Username, in.Password, c)
	if err != nil {
		//查找不到--报错
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功",
		"data": sysUser,
	})
}
