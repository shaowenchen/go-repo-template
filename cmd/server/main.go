package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/shaowenchen/go-repo-template/pkg/server"
	"github.com/shaowenchen/go-repo-template/web"
)

func init() {
	configpath := flag.String("c", "", "")
	flag.Parse()
	server.LoadConfig(*configpath)
	if _, err := server.BuildGlobalDao(); err != nil {
		println(err.Error())
	}
}

func main() {
	gin.SetMode(server.GlobalConfig.Server.RunMode)
	r := gin.Default()
	web.SetupRoutes(r)
	server.SetupV1Routes(r)
	server.SetupV1AuthRoutes(r)
	r.Run(":80")
}
