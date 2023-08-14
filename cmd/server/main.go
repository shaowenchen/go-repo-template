package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shaowenchen/go-repo-template/pkg/server"
	"github.com/shaowenchen/go-repo-template/web"
	"net/http"
	_ "net/http/pprof"
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
	web.SetupRouter(r)
	server.SetupV1Router(r)
	server.SetupV1AuthRouter(r)
	setupPprof()

	r.Run(":80")
}

func setupPprof() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", 6060), nil)
		if err != nil {
			fmt.Printf("pprof server error: %v\n", err)
		} else {
			fmt.Printf("pprof server start\n， http://localhost:6060/debug/pprof/")
		}
	}()
}
