package service

import (
	"douya/admin/entity"
	"douya/dao"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetDramaList(c *gin.Context) {
	in := new(entity.DramaListReq)
	err := c.ShouldBindQuery(&in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	list, err := dao.GetDramaList(in)
	if err != nil {
		//查找不到--报错
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "查询失败",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": list,
	})

}
