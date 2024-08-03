package model

import (
_	"time"
)

type Tag struct {
	ID uint `form:"id" json:"id" gorm:"primaryKey"` //id should be removed from form
	Name   string   `form:"name" json:"name" validate:"required" gorm:"unique"`
	//gorm.Model
}

