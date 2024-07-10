package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DB    *gorm.DB
	R     *gin.Engine
	dberr error
)

func init() {
	fmt.Println("service init")
	R = gin.Default()
	DB, dberr = gorm.Open(sqlite.Open("db/db.sqlite"))
	if dberr != nil {
		log.Fatal(dberr)
	}
}
