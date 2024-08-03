package service

//TODO:moving all hardcaoded pathes to conf variables eg. tmpls and db
import (
	"fmt"
	"log"

	//	"github.com/alaa2amz/g1/service/model"
	//"github.com/alaa2amz/g1/service/model"
	//	"github.com/alaa2amz/g1/g1dump/model"
	//"github.com/alaa2amz/g1/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	R     *gin.Engine
	dberr error
	Paths = map[string]any{}
	Index = []string{}
	StaticDir="./static"
	StaticRoute="/static"
)

func init() {
	fmt.Println("service init")

	R = gin.Default()

	R.LoadHTMLGlob("tmpl/*.tmpl")
	DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))
	if dberr != nil {
		log.Fatal(dberr)
	}
	// model.Init()
	R.GET("/",rootHandler)
	R.Static(StaticRoute,StaticDir)
}

func rootHandler(c *gin.Context) {
	log.Println("zzz")
	c.HTML(200,"root.tmpl",gin.H{"index":Index,"static":StaticRoute})
	return
}
