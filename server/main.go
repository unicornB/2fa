package main

import (
	"fmt"
	"gongniu/conf"
	"gongniu/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")
	conf.Init()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := router.NewRouter()
	r.Run(os.Getenv("PORT"))
}
