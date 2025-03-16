package model

import "time"

// sqlite 数据库 user_email_log 表
// 邮箱验证码日志
type UserEmailLog struct {
	BaseModel
	Email string `gorm:"type:varchar(100)"`
	Code  string `gorm:"type:varchar(10)"`
	Ip    string `gorm:"type:varchar(20)"`
}

func (UserEmailLog) TableName() string {
	return "user_email_log"
}

func (userEmailLog *UserEmailLog) GetTodaySendCount() int64 {
	var count int64
	DB.Model(&UserEmailLog{}).Where("created_at >= ?", time.Now().Format("2006-01-02")).Count(&count)
	return count
}

// 查询5分钟内的验证码是否正确
func (userEmailLog *UserEmailLog) GetSendCountByEmail(email, code string) int64 {
	var count int64
	DB.Model(&UserEmailLog{}).Where("email =? and code =? and created_at >=?", email, code, time.Now().Add(-5*time.Minute)).Count(&count)
	return count
}
