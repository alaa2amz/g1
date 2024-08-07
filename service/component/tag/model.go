package tag

import (
	"github.com/alaa2amz/g1/service/model"
)

var (
Path = "/tag"
	DroppedColumns        = []string{"publish_at", "afloat"}
)

type  Tag model.Tag
func Proto() (p Tag) { return }
func Protos() (p []Tag) { return }
