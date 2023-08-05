package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaowenchen/go-repo-template/pkg/server"
)

func init() {
	configpath := flag.String("c", "", "")
	flag.Parse()
	server.InitGlobalConfig(*configpath)
	if err := server.InitGlobalDao(); err != nil {
		println(err.Error())
	}
}

func main() {
	gin.SetMode(server.GlobalConfig.Server.RunMode)
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./web/dist/favicon.ico")
	r.StaticFile("/logo.png", "./web/dist/logo.png")
	r.Static("/assets", "./web/dist/assets")
	r.LoadHTMLGlob("./web/dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	v1 := r.Group("api/v1")
	{
		v1.GET("", server.Get)
		v1.GET("/version", server.Version)
	}
	v1auth := r.Group("api/v1").Use(server.AuthMiddleware())
	{
		v1auth.GET("/auth", server.GetAuth)
	}
	r.Run(":80")
}
