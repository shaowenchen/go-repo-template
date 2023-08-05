package main

import (
	"flag"
	"github.com/shaowenchen/go-repo-template/pkg/dao"
	"github.com/shaowenchen/go-repo-template/pkg/models"
	"github.com/shaowenchen/go-repo-template/pkg/server"
)

func init() {
	configpath := flag.String("c", "", "")
	flag.Parse()
	server.InitGlobalConfig(*configpath)
}

func main() {
	db, err := dao.InitDBConnect(server.GlobalConfig.DB)

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Users{})
}
