package middleware

import (
	"fmt"
	"gongniu/model"
	"gongniu/serializer"
	"gongniu/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid uint
		token := c.GetHeader("Authorization")
		if token != "" {
			user, err := utils.ParseJwt(token, os.Getenv("JWT_SECRET"))
			if err == nil {
				uid = user.ID
			}
		}
		if uid > 0 {
			admin, err := model.GetAdmin(uid)
			if err == nil {
				c.Set("admin_id", &uid)
				c.Set("admin", &admin)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录

func AuthRequired() gin.HandlerFunc {
	fmt.Println("拦截验证登录1")
	return func(c *gin.Context) {
		fmt.Println("拦截验证登录")
		if admin, _ := c.Get("admin"); admin != nil {
			if _, ok := admin.(*model.Admin); ok {
				c.Next()
				return
			} else {
				c.JSON(200, serializer.Error("未登录"))
				c.Abort()
			}
		} else {
			c.JSON(200, serializer.Error("未登录"))
			c.Abort()
		}

	}
}
