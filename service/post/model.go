package post

import (
	"github.com/alaa2amz/g1/service/model"
)

var (
	Path           string = "/post"
	DroppedColumns        = []string{"publish_at", "afloat"}
)

type Post model.Post
func Proto() (p Post) { return }
func Protos() (p []Post) { return }
