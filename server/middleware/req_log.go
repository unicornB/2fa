package middleware

import (
	"2fa.com/utils"
	"github.com/gin-gonic/gin"
)

func ReqLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.ReqLog(c)
	}
}
