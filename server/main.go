package main

import (
	"fmt"
	"os"

	"2fa.com/conf"
	"2fa.com/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")
	//utils.GenerateRSAKey(2048)
	conf.Init()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := router.NewRouter()
	r.Run(os.Getenv("PORT"))
}
