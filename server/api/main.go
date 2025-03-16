package api

import (
	"fmt"

	"2fa.com/serializer"
	"2fa.com/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//to []string 怎么初始化
	htmlContent := fmt.Sprintf(`
		<html>
		<body>
			<h2>验证码</h2>
			<p>您的验证码是：<strong>%s</strong></p>
			<p>请在 5 分钟内使用该验证码。</p>
		</body>
		</html>
	`, "123456")
	utils.SendEmailV2("hylvip2014@163.com", "2fa验证", htmlContent)
	c.JSON(200, serializer.Success("登录成功", nil))
}
