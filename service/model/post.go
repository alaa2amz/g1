package model

import (
	"time"
)

type Post struct {
	ID uint `form:"id" json:"id" gorm:"primaryKey"` //id should be removed from form
	Title   string   `form:"title" json:"title" validate:"required"`
	Content string   `form:"content" json:"content" gorm:"default:null;not null"`
	Name    *string  `form:"abstract"`
	Rate    *float64 `form:"rate"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Comments []Comment
	//gorm.Model
	//Title    string   `form:"title" json:"title" binding:"required"`
	//	PublishAt     *time.Time `form:"publish" time_format:"2006-01-02"`
	//TagID    *uint    `form:"tagid"`
	//Tag         *Tag       `form:"tag"`
}

