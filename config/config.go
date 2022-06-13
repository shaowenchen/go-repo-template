package config

import (
	"github.com/spf13/viper"
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

var Config = Options{}

func initViper(configPath string) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if configPath == "" {
		viper.SetConfigName("default")
		viper.SetConfigType("toml")
	} else {
		viper.SetConfigFile(configPath)
	}
	viper.AddConfigPath(filepath.Join(path, "."))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: %s \n", err)
	}
}

func ReadConfig(configPath string) {
	initViper(configPath)
	Config = Options{
		Gin: GinOption{viper.GetString("gin.runmode")},
	}
}

type Options struct {
	Gin GinOption
}

type GinOption struct {
	RunMode string
}
