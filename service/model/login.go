package model

import "gorm.io/gorm"

type Login struct {
	UserID uint 
	User  User 
	Name string `form:"username" gorm:"-"`
	Password string `form:"password" gorm:"-"`
	Re string `form:"re" gorm:"-"`
	TH []byte
	gorm.Model
}
