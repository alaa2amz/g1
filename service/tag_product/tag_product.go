package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"log"
	"strconv"
	"github.com/alaa2amz/g1/service/tag"
	"github.com/alaa2amz/g1/service/post"

)

var DB *gorm.DB
var Path1 string = "/tag"
var Path2 string = "/post"

type Tag tag.Tag
type Post post.Post

func Init(db *gorm.DB) {
	DB = db
}


func Register(r *gin.Engine) *gin.Engine {
	r.GET(Path, rt)
	r.POST(Path, cr)
	r.PATCH(Path + "/:id", up)
	r.DELETE(Path + "/:id", dl)
	return r
}


func cr(c *gin.Context) {
	var p Tag
	c.Bind(&p)
	r := DB.Create(&p)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func rt(c *gin.Context) {
	var p []Tag
	r := DB.Find(&p)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func up(c *gin.Context) {
	var p Tag
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	r := DB.Save(&p)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}

func dl(c *gin.Context) {
	var p Tag
	id := c.Param("id")
	c.Bind(&p)
	intid, _ := strconv.Atoi(id)
	p.ID = uint(intid)
	r := DB.Delete(&Tag{}, id)
	c.Writer.Write([]byte(fmt.Sprintf("%+v\n", r)))
	c.JSON(200, gin.H{"p": p})
}
