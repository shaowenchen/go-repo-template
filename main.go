package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/shaowenchen/go-repo-template/config"
)

var (
	GitCommit string
	BuildTime string
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
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.GET("/version", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"GitCommit": GitCommit,
			"BuildTime": BuildTime,
		})
	})
	router.Run(":80")
}
