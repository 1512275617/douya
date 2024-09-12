package dao

import (
	"context"
	"douya/components"
	"douya/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserNamePassword(username, password string, c *gin.Context) (*model.User, error) {
	db := components.GetMysqlClient(context.Background())
	data := new(model.User)
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  errors.New("用户名不能为空"),
		})
		return data, nil
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  errors.New("密码不能为空"),
		})
		return data, nil
	}
	err := db.Model(&model.User{}).Where("username=? AND password=? ", username, password).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
