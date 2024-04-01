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
	Port     int
	CharSet  string
}

func InitDBConnect(option DBOptions) (*gorm.DB, error) {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&tls=true",
		option.User, option.Password, option.Addr, option.Port, option.Name, option.CharSet,
	)
	return gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
}
