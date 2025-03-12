package conf

import (
	"gongniu/model"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()
	model.Database(os.Getenv("SQL_DSN"))
}
