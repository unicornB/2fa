package model

type User struct {
	BaseModel
	Email      string `gorm:"type:varchar(100);not null;unique" json:"email"`
	LoginIp    string `gorm:"type:varchar(20);not null;default:''" json:"login_ip"`
	Status     int    `gorm:"type:tinyint;not null;default:0" json:"status"`
	RegisterIp string `gorm:"type:varchar(20);not null;default:''" json:"register_ip"`
}

func (User) TableName() string {
	return "user"
}

func GetUser(id uint) (User, error) {
	var user User
	result := DB.Where("id = ?", id).First(&user)
	return user, result.Error
}
