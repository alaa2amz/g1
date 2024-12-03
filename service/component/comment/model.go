package comment

import (
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Path           = "/comment"
	DroppedColumns = []string{"publish_at", "afloat"}
	LeadCols       = []string{"id"}
	TrailCols      = []string{"created_at", "updated_at", "deleted_at"}
	TidyCols       = []string{}
	R              *gin.Engine
	DB             *gorm.DB
)

type Comment model.Comment

func Proto() (p Comment)    { return }
func Protos() (p []Comment) { return }
