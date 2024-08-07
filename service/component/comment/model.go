package comment

import (
	"github.com/alaa2amz/g1/service/model"
)

var (
Path = "/comment"
	DroppedColumns        = []string{"publish_at", "afloat"}
)
type Comment model.Comment
func Proto() (p Comment) { return }

func Protos() (p []Comment) { return }
