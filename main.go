package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/shaowenchen/go-repo-template/config"
)

func init() {
	configpath := flag.String("c", "", "")
	flag.Parse()
	config.ReadConfig(*configpath)
}

func main() {
	gin.SetMode(config.Config.Gin.RunMode)
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "")
	})
	router.Run(":8000")
}
