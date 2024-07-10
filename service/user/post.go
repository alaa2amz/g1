package user

import (
	"github.com/alaa2amz/g1/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var DB *gorm.DB

var Path string = "/user"
var Editions = []string{"", "api"}

type Post2 struct {
	ID       uint   `form:"id" gorm:"primaryKey"`
	Username string `form:"title"`
	Password string `form:"content"`
}

func Proto() (p Post2)    { return }
func Protos() (p []Post2) { return }

func init() {
	log.Println(Path + "init")
	if service.DB == nil {
		log.Fatal("main database not initialized")
	}
	DB = service.DB
	DB.AutoMigrate(Proto())

	if service.R == nil {
		log.Fatal("main router not initialized")
	}
	service.R = Register(service.R)
}

func Init() {
	DB.AutoMigrate(Proto())
	DB = service.DB
}

func Register(r *gin.Engine) *gin.Engine {
	for _, ed := range Editions {
		fullPath := ed + Path
		r.POST(fullPath, cr)
		r.GET(fullPath, rt)
		r.GET(fullPath+"/:id", gt)
		r.PATCH(fullPath+"/:id", up)
		r.DELETE(fullPath+"/:id", dl)
	}
	return r
}

func cr(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	DB.Create(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"data": p})
}

func rt(c *gin.Context) {
	//var p []Post
	p := Protos()
	DB.Find(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"data": p})
}

func gt(c *gin.Context) {
	var p Post2
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	DB.First(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func up(c *gin.Context) {
	var p Post2
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	DB.Save(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func dl(c *gin.Context) {
	var p Post2
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	DB.Delete(&Post2{}, id)
	//r := DB.Delete(&Post{}, id)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}
