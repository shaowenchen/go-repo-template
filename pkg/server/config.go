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

func LoadConfig(configPath string) {
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
	err = viper.Unmarshal(GlobalConfig)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v \n", err)
	}
}

func BuildGlobalDao() (*gorm.DB, error) {
	GlobalDB, err := dao.InitDBConnect(GlobalConfig.DB)
	return GlobalDB, err
}
