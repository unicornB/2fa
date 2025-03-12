package middleware

import (
	"fmt"
	"gongniu/model"
	"gongniu/serializer"
	"gongniu/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid uint
		token := c.GetHeader("Authorization")
		fmt.Println("token:" + token)
		if token != "" {
			user, err := utils.ParseJwt(token, os.Getenv("JWT_SECRET_USER"))
			if err == nil {
				uid = user.ID
			}
		}

		if uid > 0 {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user_id", &uid)
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

func AuthUserRequired() gin.HandlerFunc {
	fmt.Println("拦截验证登录1")
	return func(c *gin.Context) {
		fmt.Println("拦截验证登录")
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			} else {
				c.JSON(200, serializer.CheckLogin())
				c.Abort()
			}
		} else {
			c.JSON(200, serializer.CheckLogin())
			c.Abort()
		}

	}
}
