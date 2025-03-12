package model

type BizData struct {
	Id         int    `gorm:"primary_key"`
	BizKey     string `gorm:"type:varchar(50);not null;unique"`
	BizData    string `gorm:"type:text"`
	Desc       string `gorm:"type:varchar(100)"`
	StartMonth string `gorm:"type:int(3);remark:开始月份"`
}

func (BizData) TableName() string {
	return "biz_data"
}
