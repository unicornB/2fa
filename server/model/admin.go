package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id       int    `gorm:"primary_key"`
	Username string `gorm:"type:varchar(30);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}

func (Admin) TableName() string {
	return "admin"
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	admin.SetPassword(password)
	return err == nil
}
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.Password = string(bytes)
	fmt.Println("密码:" + admin.Password)
	return nil
}
func GetAdmin(ID interface{}) (Admin, error) {
	var admin Admin
	result := DB.First(&admin, ID)
	return admin, result.Error
}
func (admin *Admin) InitAdmin() {
	//判断是否有admin
	var count int64
	DB.Model(&Admin{}).Count(&count)
	if count == 0 {
		//初始化管理员
		bytes, _ := bcrypt.GenerateFromPassword([]byte("123456"), PassWordCost)
		_ = DB.Create(&Admin{
			Username: "admin",
			Password: string(bytes),
		}).Error
	}
}
