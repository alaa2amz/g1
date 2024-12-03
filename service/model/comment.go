package model

import (
	"gorm.io/gorm"
)

/*
var (
	Path           string = "/post"
	DroppedColumns        = []string{"publish_at", "afloat"}
)
*/
// type Content comment.Comment
type Comment struct {
	gorm.Model
	//UserID uint
	//User User
	//PostID  *uint64 `gorm:"default:null"`
	//Post    *Post
	Title   string `form:"ctitle" json:"title" validate:"required"`
	Content string `form:"ccontent" json:"content" gorm:"not null"`
}

//type Tag tag.Tag

//func Proto() (p Post) { return }

//func Protos() (p []Post) { return }
