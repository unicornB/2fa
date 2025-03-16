package service

import (
	"fmt"
	"os"
	"strconv"

	"2fa.com/model"
	"2fa.com/serializer"
	"2fa.com/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	Email string `form:"username" json:"username" binding:"min=3,max=30"`
	Code  string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

// 邮箱验证码登录，首先判断是否存在该用户，不存在则创建，存在则直接登录
func (service *UserService) Login() serializer.Response {
	email, err := utils.RSA_Decrypt(service.Email)
	if err != nil {
		return serializer.Error("对不起，不支持你的邮箱")
	}
	code, err := utils.RSA_Decrypt(service.Code)
	if err != nil {
		return serializer.Error("对不起，不支持你的验证码")
	}
	//判断验证码是否正确，并且在5分钟内有效
	var userEmailLog model.UserEmailLog
	count := userEmailLog.GetSendCountByEmail(email, code)
	if count == 0 {
		return serializer.Error("验证码错误")
	}
	//判断是否存在该用户
	var user model.User
	user.Email = email
	DB := model.DB
	result := DB.Where("email =?", email).First(&user)
	if result.Error != nil {
		//不存在则创建
		user := model.User{
			Email: email,
		}
		DB.Create(&user)
	}
	//生成token
	token, err := utils.GenerateJWT(email, os.Getenv("JWT_SECRET_USER"), user.Id)
	if err != nil {
		return serializer.Error("对不起，系统繁忙，请稍后再试")
	}
	return serializer.Success("登录成功", token)
}

type UserSendMailService struct {
	Email string `form:"email" json:"email" binding:"required"`
}

// 发送邮箱验证码
func (service *UserSendMailService) SendEmail(c *gin.Context) serializer.Response {
	email, err := utils.RSA_Decrypt(service.Email)
	if err != nil {
		return serializer.Error("对不起，不支持你的邮箱")
	}
	fmt.Println("email:" + email)
	//获取今日发送的验证码次数
	var emailLog model.UserEmailLog
	emailLog.Email = email
	count := emailLog.GetTodaySendCount()
	emailLimit := os.Getenv("TODAY_SEND_EMAIL_COUNT")
	//字符串转int64
	limit, err := strconv.ParseInt(emailLimit, 10, 64)
	if err != nil {
		return serializer.Error("系统配置错误")
	}
	if count >= limit {
		return serializer.Error("今日发送次数已达上限")
	}
	code := utils.GenerateRandomNumber()
	//发送邮件
	htmlContent := fmt.Sprintf(`
		<html>
		<head>
			<style>
				body { font-family: Arial, sans-serif; background-color: #f5f5f5; padding: 20px; }
				.container { max-width: 600px; margin: 0 auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
				h2 { color: #333; margin-bottom: 20px; }
				.code { font-size: 24px; font-weight: bold; color: #007bff; padding: 10px 20px; background: #e9f5ff; border-radius: 4px; display: inline-block; margin: 10px 0; }
				.note { color: #666; margin-top: 20px; font-size: 14px; }
				.footer { margin-top: 30px; padding-top: 20px; border-top: 1px solid #eee; color: #999; font-size: 12px; }
			</style>
		</head>
		<body>
			<div class="container">
				<h2>验证码</h2>
				<p>您的验证码是：</p>
				<div class="code">%s</div>
				<p class="note">请在 5 分钟内使用该验证码。</p>
				<div class="footer">
					<p>如果您没有请求此验证码，请忽略此邮件。</p>
				</div>
			</div>
		</body>
		</html>
	`, code)
	sendErr := utils.SendEmailV2(email, os.Getenv("EMAIL_FROM"), htmlContent, os.Getenv("EMAIL_FROM_NAME"))
	if sendErr != nil {
		return serializer.Error("发送失败")
	}
	//保存验证码
	userEmailLog := model.UserEmailLog{
		Email: email,
		Code:  code,
		Ip:    c.ClientIP(),
	}
	model.DB.Create(&userEmailLog)
	return serializer.Success("发送成功", nil)
}
