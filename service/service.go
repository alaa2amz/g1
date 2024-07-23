package service

//TODO:moving all hardcaoded pathes to conf variables eg. tmpls and db
import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	R     *gin.Engine
	dberr error
)

func init() {
	fmt.Println("service init")

	R = gin.Default()

	R.LoadHTMLGlob("service/tmpl/*.tmpl")
	DB, dberr = gorm.Open(sqlite.Open("db/db.sqlite"))
	if dberr != nil {
		log.Fatal(dberr)
	}
}
