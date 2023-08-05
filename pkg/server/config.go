package server

import (
	"fmt"
	"github.com/shaowenchen/go-repo-template/pkg/dao"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

var GlobalConfig = &ConfigOptions{}
var GlobalDB = &gorm.DB{}

var (
	GitCommit string
	BuildTime string
)

type ConfigOptions struct {
	Server ServerOptions
	DB     dao.DBOptions
}

type ServerOptions struct {
	RunMode string
}

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

func InitGlobalConfig(configPath string) *ConfigOptions {
	initViper(configPath)
	viper.SetDefault("server.runmode", "debug")
	GlobalConfig.Server.RunMode = viper.GetString("server.runmode")
	GlobalConfig.DB.User = viper.GetString("db.user")
	GlobalConfig.DB.Password = viper.GetString("db.password")
	GlobalConfig.DB.Addr = viper.GetString("db.addr")
	GlobalConfig.DB.Name = viper.GetString("db.name")
	return GlobalConfig
}

func InitGlobalDao() (err error) {
	GlobalDB, err = dao.InitDBConnect(GlobalConfig.DB)
	return
}
