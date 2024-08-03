package model

import "gorm.io/gorm"

type User struct {
	Name string `form:"username" json:"username" validate:"required" gorm:"unique"`
	Password   string `form:"password" gorm:"-"`
	Confirm   string `form:"confirm" gorm:"-"`
	PH   []byte
	gorm.Model
}
