package post

import (
	"github.com/alaa2amz/g1/service/model"
)

var (
	Path           string = "/post"
	DroppedColumns        = []string{"publish_at", "afloat"}
)
/*
// type Content comment.Comment
type Post struct {
	ID uint `form:"id" json:"id" gorm:"primaryKey"` //id should be removed from form
	//Title    string   `form:"title" json:"title" binding:"required"`
	Title   string   `form:"title" json:"title" validate:"required"`
	Content string   `form:"content" json:"content" gorm:"default:null;not null"`
	Name    *string  `form:"abstract"`
	Rate    *float64 `form:"rate"`
	//TagID    *uint    `form:"tagid"`
	//Tag         *Tag       `form:"tag"`
	//	PublishAt     *time.Time `form:"publish" time_format:"2006-01-02"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	//gorm.Model
	//	Comments []comment.Comment
}
*/
//type Tag tag.Tag
type Post model.Post
func Proto() (p Post) { return }

func Protos() (p []Post) { return }
