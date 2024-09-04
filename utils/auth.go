package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"
)

const (
	HeaderAuthorization = "Authorization"
)

type MyClaims struct {
	AppId int `json:"app_id"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("123456")

// GenerateToken 生成JWT
func GenerateToken() (string, error) {
	c := MyClaims{
		1,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "douya",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tmp, err := token.SignedString(mySecret)
	return tmp, err
}

func ParseToken(tokenSring string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenSring, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if Claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return Claims, nil
	}
	return nil, errors.New("invalid token")
}

// Handler ...
// @loginRequired 为true时则接口必须登录，否则直接返回鉴权错误。
func Handler(loginRequired bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		if !loginRequired {
			c.Next()
			return
		}
		token := c.Request.Header.Get(HeaderAuthorization)
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]any{
				"code": strconv.Itoa(http.StatusUnauthorized),
				"msg":  "token不存在",
			})

			c.Abort()
			return
		}
		_, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]any{
				"code": strconv.Itoa(http.StatusUnauthorized),
				"msg":  "校验失败",
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
