package wig

import (
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
Path = "/wig"
	DroppedColumns = []string{"publish_at", "afloat"}
	LeadCols       = []string{"id"}
	TrailCols      = []string{"created_at", "updated_at", "deleted_at"}
	TidyCols       = []string{}
	R              *gin.Engine
	DB             *gorm.DB
)

type  Wig model.Wig

func Proto() (p Wig)    { return }
func Protos() (p []Wig) { return }
