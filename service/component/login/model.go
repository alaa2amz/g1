package login

import (
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	Path           = "/login"
	DroppedColumns = []string{"publish_at", "afloat"}
	LeadCols       = []string{"id"}
	TrailCols      = []string{"created_at", "updated_at", "deleted_at"}
	TidyCols       = []string{}
	R              *gin.Engine
	DB             *gorm.DB
)

//type  Login model.Login

func Proto() (p model.Login)    { return }
func Protos() (p []model.Login) { return }
