package post

import (
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Path           string = "/post"
	DroppedColumns        = []string{"publish_at", "afloat"}
	LeadCols              = []string{"id"}
	TrailCols             = []string{"created_at", "updated_at", "deleted_at"}
	TidyCols              = []string{}
	R              *gin.Engine
	DB             *gorm.DB
)

type Post model.Post

func Proto() (p Post)    { return }
func Protos() (p []Post) { return }
