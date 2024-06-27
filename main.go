package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"

	"github.com/alaa2amz/g1/service/post"
	"github.com/alaa2amz/g1/service/tag"
)

var r  *gin.Engine
var DB *gorm.DB
var dberr error

func init() {

	r = gin.Default()
	DB, dberr = gorm.Open(sqlite.Open("db.sqlite"))
	if dberr != nil {
		log.Fatal(dberr)
	}

	r =post.Register(r)
	post.Init(DB)
	r =tag.Register(r)
	tag.Init(DB)
}
func main() {

	fmt.Println("Hello, Web!")
	r.Run()
}

