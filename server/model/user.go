package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(100);not null;unique" json:"username"`
	Password string `gorm:"type:varchar(255);not null;" json:"password"`
}

func (User) TableName() string {
	return "user"
}
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	user.SetPassword(password)
	return err == nil
}
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	fmt.Println("密码:" + user.Password)
	return nil
}
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}
func (admin *User) InitUser() {
	//判断是否有admin
	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		//初始化管理员
		bytes, _ := bcrypt.GenerateFromPassword([]byte("123456"), PassWordCost)
		_ = DB.Create(&User{
			Username: "admin",
			Password: string(bytes),
		}).Error
	}
}
