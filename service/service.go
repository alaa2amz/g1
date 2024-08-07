package service

//TODO:moving all hardcaoded pathes to conf variables eg. tmpls and db
import (
	"fmt"
	"log"

	//	"os"

	//	"github.com/alaa2amz/g1/service/model"
//	"github.com/alaa2amz/g1/service/model"
//	"github.com/alaa2amz/g1/g1dump/model"
//	"github.com/alaa2amz/g1/service/model"
//	"os"
	//"github.com/alaa2amz/g1/model"
	"github.com/gin-gonic/gin"
	//	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
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
	
	//sqlite//DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))
	//DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))
	
  dsn := "g1gorm:password12346@tcp(127.0.0.1:3306)/g1?charset=utf8mb4&parseTime=True&loc=Local"
DB, dberr = gorm.Open(mysql.Open(dsn), &gorm.Config{})



//logger.New()


//DB, dberr = gorm.Open(mysql.Open(dsn), &gorm.Config{
    //Logger: logger.Default.LogMode(logger.Info),
 //   Logger: fileLogger.LogMode(logger.Info),
//})

	if dberr != nil {
		log.Fatal(dberr)
	}
	DB=DB.Debug()
	/*DB.AutoMigrate(&model.User{},
			&model.Login{},
			&model.Post{},
			&model.Tag{},
			&model.Comment{},
		)*/
	// model.Init()
	R.GET("/",rootHandler)
	R.Static(StaticRoute,StaticDir)
}

func rootHandler(c *gin.Context) {
	log.Println("zzz")
	c.HTML(200,"root.tmpl",gin.H{"index":Index,"static":StaticRoute})
	return
}
