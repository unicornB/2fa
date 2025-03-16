package conf

import (
	"fmt"
	"os"

	"2fa.com/model"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()
	initFile()
	InitLogger()
	model.Database("./db/2fa.db")
}
func initFile() {
	fmt.Println("initFile")
	//判断是否存在./db/2fa.db文件,不存在则创建
	if _, err := os.Stat("./db/2fa.db"); os.IsNotExist(err) {
		os.Mkdir("./db", 0755)
		os.Create("./db/2fa.db")
		fmt.Println("2fa.db不存在,创建成功")
	}
	//判断是否存在./logs/app.log文件,不存在则创建
	if _, err := os.Stat("./logs/app.log"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0755)
		os.Create("./logs/app.log")
		fmt.Println("app.log不存在,创建成功")
	}
}
