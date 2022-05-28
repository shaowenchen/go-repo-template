package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vinqi/vqchat/mp"
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
	mp.ReadConfig()
	gin.SetMode(mp.Config.RunMode)
}

var VERSION string = "latest"

func main() {
	go mp.BeatVqdata()
	router := gin.Default()
	router.POST("/", mp.Handler)
	router.GET("/", func(context *gin.Context) {
		context.String(200, VERSION)
	})
	router.Run(":8000")
}
