package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	//UserID   *uint ` gorm:"default:null"`

	//User     *User
	Title   string   `form:"title" json:"title" validate:"required"`
	Content string   `form:"content" json:"content" gorm:"not null"`
	Name    *string  `form:"abstract"`
	Rate    *float64 `form:"rate"`
	//Comments []Comment
}

//Content string   `form:"content" json:"content" gorm:"default:null;not null"`
//Title    string   `form:"title" json:"title" binding:"required"`
//	PublishAt     *time.Time `form:"publish" time_format:"2006-01-02"`
//TagID    *uint    `form:"tagid"`
//Tag         *Tag       `form:"tag"`
