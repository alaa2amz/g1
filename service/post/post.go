// Alaa Ahmed Zakariya.
// Post ressource.
// Example component.
// a Social text.
// holds model, web , api ,test.
// TODO:generic joiner
// TODO:edit and new points
package post

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/alaa2amz/g1/service"
	"github.com/alaa2amz/g1/service/tag"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	Path         string = "/post"
	Editions            = []string{"", "api"}
	TemplatePath        = "post/template"
)

type Reply struct {
	StatusCode int
	Data       any
	Error      error
	Template   string
}

type QueryTerm struct {
	Or       bool
	Column   string
	Relation string
	//Values   string
	Values []string
}

var relations = map[string]string{
	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
	"in": "IN",
	"ni": "NOT IN",
	"co": "LIKE",
}

type Post struct {
	ID       uint     `form:"id" gorm:"primaryKey"`
	Title    string   `form:"title" binding:"required"`
	Descript *string  `form:"descript"`
	Content  string   `form:"content" gorm:"default:null;not null"`
	Abstract *string  `form:"abstract"`
	Afloat   float64  `form:"afloat"`
	Rate     *float64 `form:"rate"`
	//TagID    *uint    `form:"tagid"`
	//Tag      *Tag     `form:"tagname"`
	//Tag         *Tag       `form:"tag"`
	PublishAt     *time.Time `form:"publish"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	//gorm.Model

}

type Tag tag.Tag

func Proto() (p Post) { return }

func Protos() (p []Post) { return }

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
	//TODO:log only
	DB.AutoMigrate(Proto())
	DB = service.DB
}

// Register
func Register(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("**/post/template/*.tmpl")

	for _, ed := range Editions {
		fullPath := ed + Path
		//end points in CRUD order
		// cr -> create
		// rt -> retrive
		// gt -> get by ID
		// up -> update by id
		// dl -> delete by id
		///// helpers
		// nw -> new form or empty json
		// ed -> edit form or json by id
		// jg -> join generic other resoursec
		r.POST(fullPath, cr)
		r.GET(fullPath, rt)
		r.GET(fullPath+"/:id", gt)
		r.PUT(fullPath+"/:id", up)
		r.DELETE(fullPath+"/:id", dl)

		//r.GET(fullPath+"/test", tst)
		//r.GET(fullPath+"/new", nw)
		//r.GET(fullPath+"/:id/edit", ed)
	}
	return r
}

// TODO: c.send(reply)
func send(r *Reply, c *gin.Context) {
	errMsg := ""
	if r.Error != nil {
		errMsg = r.Error.Error()
	}
	if ok := strings.Contains(c.Request.URL.Path, "/api/"); ok {
		//handeling data and errors
		c.JSON(r.StatusCode, gin.H{"data": r.Data, "error": errMsg})
		return
	} else {
		c.HTML(r.StatusCode, r.Template, gin.H{"data": r.Data, "error": errMsg})
		return
	}
}

func parseQueryString(qs map[string][]string) []QueryTerm {
	terms := []QueryTerm{}
	for key, values := range qs {
		for _, value := range values {
			term := QueryTerm{}
			term.Column = key
			vals := strings.Split(value, "__")
			if len(vals) == 1 {
				term.Values = append(term.Values, vals[0])
				term.Relation = "="
				terms = append(terms, term)
				continue

			} else if vals[0][0:2] == "or" {
				term.Or = true
				vals[0] = vals[0][2:]
			}
			term.Relation = relations[vals[0]]
			term.Values = append(term.Values, vals[1:]...)
			terms = append(terms, term)
		}
	}
	return terms
}
func tst(c *gin.Context) {
	r := &Reply{}
	send(r, c)
}

// cr Create Handler
func cr(c *gin.Context) {
	r := &Reply{}
	p := Proto()

	err := c.ShouldBind(&p)
	if err != nil {
		r.StatusCode = 400
		r.Error = err
		r.Template = "error.tmpl"
		send(r, c)
		return
	}

	result := DB.Create(&p)
	if result.Error != nil {
		r.StatusCode = 400
		r.Error = result.Error
		r.Template = "error.tmpl"
		send(r, c)
		return
	}
	r.StatusCode = 200
	r.Data = p
	send(r, c)
	return
}

func rt(c *gin.Context) {
	qs := c.Request.URL.Query()
	terms := parseQueryString(qs)
	log.Printf("%+v\n", terms)
	//p := Proto()
	ps := Protos()
	QDB := DB.Session(&gorm.Session{})
	QDB.Model(&ps)
	for _, term := range terms {
		if term.Relation == "LIKE" {
			term.Values[0] = "%" + term.Values[0] + "%"
		}
		queryString := fmt.Sprintf("%s %s ?", term.Column, term.Relation)
		log.Println(queryString)
		log.Printf("%T ----%v\n", term.Values, term.Values)
		if !term.Or {
			//QDB = QDB.Where(queryString, term.Values[0]).Where("id > ?",100)
			QDB = QDB.Where(queryString, term.Values[0])
		} else {

			QDB = QDB.Or(queryString, term.Values[0])
		}
	}
	log.Println(QDB.ToSQL(func(q *gorm.DB) *gorm.DB { return q.Find(&ps) }))
	QDB.Find(&ps)
	c.JSON(200, gin.H{"data": ps})
}

func gt(c *gin.Context) {
	p := Proto()
	m := map[string]interface{}{}
	id := c.Param("id")
	result := DB.Model(&p).First(&m, id)
	if result.Error != nil {
		//c.JSON(400, gin.H{"error": result.Error.Error()})
		r := &Reply{400, nil, result.Error, "error.tmpl"}
		send(r, c)
		return
	}
	r := &Reply{200, &m, nil, "show.tmpl"}
	send(r, c)

}

func up(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	p.ID = uint(uintID)
	DB.Save(&p)
	c.JSON(200, gin.H{"data": p})
}

func up2(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	id := c.Param("id")
	DB.Model(&p).Where("ID=?", id).Updates(&p)
	c.JSON(200, gin.H{"p": p})
}

func dl(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	id := c.Param("id")
	c.Bind(&p)
	DB.Delete(&p, id)
	c.JSON(200, gin.H{"data": p})
}

//result := DB.First(&p, id)
//c.JSON(200, gin.H{"data": p})
//result := DB.Model(&p).Where("ID = ?", id).First(&m)
//c.JSON(200, gin.H{"data": p})
