package model

import (
	"time"
)
/*
var (
	Path           string = "/post"
	DroppedColumns        = []string{"publish_at", "afloat"}
)
*/
// type Content comment.Comment
type Comment struct {
	ID uint `form:"id" json:"id" gorm:"primaryKey"` //id should be removed from form
	//Title    string   `form:"title" json:"title" binding:"required"`
	Title   string   `form:"ctitle" json:"title" validate:"required"`
	Content string   `form:"ccontent" json:"content" gorm:"default:null;not null"`
	Name    *string  `form:"abstract"`
	Rate    *float64 `form:"rate"`
	//TagID    *uint    `form:"tagid"`
	//Tag         *Tag       `form:"tag"`
	//	PublishAt     *time.Time `form:"publish" time_format:"2006-01-02"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	//gorm.Model
	PostID *uint64 `gorm:"default:null"`
	Post *Post
	Rrr string

}

//type Tag tag.Tag

//func Proto() (p Post) { return }

//func Protos() (p []Post) { return }
