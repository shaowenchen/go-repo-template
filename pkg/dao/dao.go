package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBOptions struct {
	User     string
	Password string
	Addr     string
	Name     string
}

func InitDBConnect(option DBOptions) (*gorm.DB, error) {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		option.User, option.Password, option.Addr, option.Name,
	)
	return gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
}
