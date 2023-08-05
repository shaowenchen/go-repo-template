package models

import (
	"time"
)

type Users struct {
	UserID    string    `gorm:"column:userid" json:"userid"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"-"`
	Status    int       `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:createdat" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedat" json:"-"`
}

func (Users) TableName() string {
	return "users"
}

const (
	_ = iota
	UsersStatusOK
	UsersStatusBlock
	UsersStatusDelete
)
