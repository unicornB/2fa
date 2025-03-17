package middleware

import (
	"fmt"
	"os"

	"2fa.com/model"
	"2fa.com/serializer"
	"2fa.com/utils"

	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	fmt.Println("CurrentUser")
	return func(c *gin.Context) {
		var uid uint
		//"Authorization": "Bearer $token",
		//token := c.GetHeader("Authorization")
		token := c.GetHeader("Authorization")
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
	return func(c *gin.Context) {
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
