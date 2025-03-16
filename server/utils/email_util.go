package utils

import (
	"fmt"
	"mime"
	"net/smtp"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendEmail(to []string, subject, body string) error {
	// 配置 Gmail 服务器信息
	smtpServer := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")
	// 邮件主题和内容
	message := []byte("Subject: " + subject + "\n\n" + body)
	// 创建认证信息
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"), smtpServer)
	// 连接到 SMTP 服务器
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, os.Getenv("EMAIL_USER"), to, message)
	if err != nil {
		fmt.Println(err)
		logrus.Error(err)
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

func SendEmailV2(to string, subject, body, title string) error {
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", os.Getenv("EMAIL_USER"), mime.QEncoding.Encode("UTF-8", title))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	smtpPort := os.Getenv("EMAIL_PORT")
	port, _ := strconv.Atoi(smtpPort)
	n := gomail.NewDialer(os.Getenv("EMAIL_HOST"), port, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"))
	if err := n.DialAndSend(msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
