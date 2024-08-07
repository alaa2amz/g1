package service

//TODO:moving all hardcaoded pathes to conf variables eg. tmpls and db
import (
	"fmt"
	"log"


	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
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
	
	
	dsn := "alaazak:0100ZakAD@/alaazak_g1?charset=utf8mb4&parseTime=True"
  	//dsn := "alaazak:0100ZakAD@tcp(127.0.0.1:3306)/alaazak_g1?charset=utf8mb4&parseTime=True&loc=Local"
	//AD//dsn := "alaazak:0100ZakAD@tcp(mysql-alaazak.alwaysdata.net:3306)/alaazak_g1?charset=utf8mb4&parseTime=True"
	
	//sqlite//DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))
	//DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))
	DB, dberr = gorm.Open(mysql.Open(dsn), &gorm.Config{})



	if dberr != nil {
		log.Fatal(dberr)
	}
	//DB=DB.Debug()
	R.GET("/",rootHandler)
	R.Static(StaticRoute,StaticDir)
}

func rootHandler(c *gin.Context) {
	c.HTML(200,"root.tmpl",gin.H{"index":Index,"static":StaticRoute})
	return
}

/*//LOGGER
logger.New()
DB, dberr = gorm.Open(mysql.Open(dsn), &gorm.Config{
Logger: logger.Default.LogMode(logger.Info),
Logger: fileLogger.LogMode(logger.Info),
})
*/

/*DB.AutoMigrate(&model.User{},
&model.Login{},
&model.Post{},
&model.Tag{},
&model.Comment{},
)*/
// model.Init()
