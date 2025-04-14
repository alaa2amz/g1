package model

import "gorm.io/gorm"

type Egg struct {
	gorm.Model
	Name string `form:"name" json:"name" validate:"required" gorm:"unique"`
	rate string `  validate:"required" gorm:"unique"`
}
