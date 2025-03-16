package router

import (
	"os"

	"2fa.com/api"
	"2fa.com/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//日志请求
	r.Use(middleware.ReqLogMiddleware())
	//签名验证
	r.Use(middleware.SignatureMiddleware(os.Getenv("SIGN_SECRET_KEY")))
	r.GET("/", api.Login)
	//分组 并添加中间件
	user := r.Group("/api/user")
	user.POST("/login", api.UserLogin)
	user.POST("/send_email_code", api.UserSendEmailCode)
	user.Use(middleware.CurrentUser()).
		Use(middleware.AuthUserRequired())
	{
		user.GET("/getMe", api.UserGetMe)
	}
	return r
}
