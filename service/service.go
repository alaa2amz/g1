package service

//TODO:moving all hardcaoded pathes to conf variables eg. tmpls and db
import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/inflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	//"gorm.io/driver/sqlite"
)

var (
	DB          *gorm.DB
	R           *gin.Engine
	dberr       error
	Paths       = map[string]any{}
	Index       = []string{}
	Childs      = map[string][]string{}
	StaticDir   = "./static"
	StaticRoute = "/static"
)

func init() {
	fmt.Println("service init", Index)

	R = gin.Default()

	/*AD
	dsn := `alaazak:0100ZakAD@tcp(mysql-alaazak.alwaysdata.net:3306)/alaazak_g1?`
		+`charset=utf8mb4&parseTime=True`
	*/

	//sqlite
	//DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))

	//local mysql
	dsn := "alaazak:0100ZakAD@/alaazak_g1?charset=utf8mb4&parseTime=True"

	//prepairing DB
	DB, dberr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dberr != nil {
		log.Fatal(dberr)
	}

	//DB=DB.Debug()

	R.GET("/", rootHandler)
	R.Static(StaticRoute, StaticDir)
	//TODO: handling no route
}

func rootHandler(c *gin.Context) {
	c.HTML(200, "root.tmpl", gin.H{"index": Index, "static": StaticRoute})
	return
}

func PostMigrate() {
	tables, err := DB.Migrator().GetTables()
	if err != nil {
		log.Panic(err)
	}

	for _, table := range tables {
		cols, err := DB.Migrator().ColumnTypes(table)
		if err != nil {
			log.Panic(err)
		}
		
		for _, col := range cols {
			if cut, ok := strings.CutSuffix(col.Name(), "_id"); ok {
				colPlural := inflection.Plural(cut)
				Childs[colPlural] = append(Childs[colPlural], table)
			}

		}
	}
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
/////////////////////
///////	R.LoadHTMLGlob("tmpl/*.tmpl")
//dsn := "alaazak:0100ZakAD@tcp(127.0.0.1:3306)/alaazak_g1?charset=utf8mb4&parseTime=True&loc=Local"
//fmt.Printf("%+v", col)
//Childs[table]=append(Childs[table],inflection.Plural(cut))
