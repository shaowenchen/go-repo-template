package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shaowenchen/go-repo-template/config"
	"github.com/spf13/viper"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(filepath.Join(path, "conf"))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("run")
	viper.SetConfigType("toml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Sprintf("fatal error config file: %s \n", err)
	}
	config.ReadConfig()
	gin.SetMode(config.Config.RunMode)
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "")
	})
	router.Run(":8000")
}
