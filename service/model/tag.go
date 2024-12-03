package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `form:"name" json:"name" validate:"required" gorm:"unique"`
	rate string `form:"name" json:"name" validate:"required" gorm:"unique"`
}
