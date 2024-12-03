package user

import (
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Path           = "/user"
	DroppedColumns = []string{"publish_at", "afloat"}
	LeadCols       = []string{"id"}
	TrailCols      = []string{"created_at", "updated_at", "deleted_at"}
	TidyCols       = []string{}
	R              *gin.Engine
	DB             *gorm.DB
)

//type  User model.User

func Proto() (p model.User)    { return }
func Protos() (p []model.User) { return }
