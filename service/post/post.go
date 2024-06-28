package post

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/alaa2amz/g1/service"
	"github.com/alaa2amz/g1/service/tag"
	"strconv"
)

var DB *gorm.DB
var Path string = "/post"

type Post struct {
	ID      uint    `form:"id" gorm:"primaryKey"`
	Title   string  `form:"title"`
	Content string  `form:"content"`
	Afloat  float64 `form:"afloat"`
	TagID   uint    `form:"tag"`
	Tag     Tag     `form:"tag"`
}

func Proto() (p Post) {return}

type Tag tag.Tag
func init() {
	log.Println(Path + "init")
	if service.DB == nil {
		log.Fatal("main database not initialized")
	}
	DB = service.DB
	DB.AutoMigrate(&Post{})

	if service.R == nil {
		log.Fatal("main router not initialized")
	}
	service.R = Register(service.R)
}

func Init() {
	DB.AutoMigrate(&Post{})
	DB = service.DB
}

func Register(r *gin.Engine) *gin.Engine {
	r.POST(Path, cr)
	r.GET(Path, rt)
	r.GET(Path+"/:id", gt)
	r.PATCH(Path+"/:id", up)
	r.DELETE(Path+"/:id", dl)
	return r
}

func cr(c *gin.Context) {
	//var p Post
	p := Proto()
	c.Bind(&p)
	DB.Create(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"data": p})
}

func rt(c *gin.Context) {
	var p []Post
	DB.Find(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"data": p})
}

func gt(c *gin.Context) {
	var p Post
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	DB.First(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func up(c *gin.Context) {
	var p Post
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	DB.Save(&p)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func dl(c *gin.Context) {
	var p Post
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	DB.Delete(&Post{}, id)
	//r := DB.Delete(&Post{}, id)
	//c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}
