package main

import (
	"github.com/shaowenchen/go-repo-template/pkg/dao"
	"github.com/shaowenchen/go-repo-template/pkg/models"
	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./pkg/dao",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(models.Users{})

	g.ApplyInterface(func(dao.UsersInterface) {}, models.Users{})

	g.Execute()
}
