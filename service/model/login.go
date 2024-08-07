package model

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	UserID   uint
	User     User
	Name     string `form:"username" gorm:"-"`
	Password string `form:"password" gorm:"-"`
	Re       string `form:"re" gorm:"-"` //referer
	TH       []byte
}
