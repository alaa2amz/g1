package post

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"log"
	"strconv"
)

var DB *gorm.DB
var Path string = "/post"

type Post struct {
	ID      uint    `form:"id" gorm:"primaryKey"`
	Title   string  `form:"title"`
	Content string  `form:"content"`
	Afloat  float64 `form:"afloat"`
}

func Init(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&Post{})
}


func Register(r *gin.Engine) *gin.Engine {
	r.GET("/", home)
	r.GET(Path, rt)
	r.POST(Path, cr)
	r.PATCH(Path + "/:id", up)
	r.DELETE(Path + "/:id", dl)
	return r
}

func home(c *gin.Context) {
	c.String(200, "Marhabah")
}

func cr(c *gin.Context) {
	var p Post
	c.Bind(&p)
	r := DB.Create(&p)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func rt(c *gin.Context) {
	var p []Post
	r := DB.Find(&p)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func up(c *gin.Context) {
	var p Post
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	r := DB.Save(&p)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func dl(c *gin.Context) {
	var p Post
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	r := DB.Delete(&Post{}, id)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}
