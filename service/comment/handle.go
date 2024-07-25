// Alaa Ahmed Zakariya.
// Post ressource.
// Example component.
// a Social text.
// holds model, web , api ,test.
// TODO:generic joiner
// TODO:edit and new points
// TODO:validate basic data types on create
// TODO:intensive error handling and testing
package comment

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	//	"github.com/alaa2amz/g1/mw"

	//	"github.com/mitchellh/mapstructure"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	R  *gin.Engine
	DB *gorm.DB
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
		//TODO:handle templates errors and unfound templates or empty templates
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
			vals := strings.Split(value, "~")
			if len(vals) == 1 {
				term.Values = append(term.Values, vals[0])
				term.Relation = "="
				terms = append(terms, term)
				continue

			} else if vals[0][0:2] == "or" {
				term.Or = true
				vals[0] = vals[0][2:]
			}
			if v, ok := relations[vals[0]]; ok {
				term.Relation = v
				term.Values = append(term.Values, vals[1:]...)
			} else {
				term.Relation = "="
			}
			switch vals[0] {
			case "co":
				term.Values[0] = "%" + term.Values[0] + "%"
			default:
			}
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
	//

	//TODO:use send rc
	c.Set("message", "updated")
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}

}

// rt retrieve records C{R}UD
func rt(c *gin.Context) {
	//TODO:sort columns and rows
	//TODO:smart sort
	//TODO:pagination
	ps := Protos()              //model {p}rototype{s}
	rm := []map[string]any{}    //{r}esults {m}ap
	r := &Reply{}               //response reply
	qs := c.Request.URL.Query() //query string map
	terms := parseQueryString(qs)
	log.Printf("terms: %+v\n", terms)  //debug
	QDB := DB.Session(&gorm.Session{}) //start session QDB
	QDB.Model(&ps)
	for _, term := range terms {
		queryString := fmt.Sprintf("%s %s ?", term.Column, term.Relation)
		if !term.Or {
			QDB = QDB.Where(queryString, term.Values[0])
		} else {

			QDB = QDB.Or(queryString, term.Values[0])
		}
	}
	log.Println(QDB.ToSQL(func(q *gorm.DB) *gorm.DB { return q.Find(&ps) })) //debug
	result := QDB.Model(&ps).Order("id desc").Find(&rm)
	if result.Error != nil {
		r.StatusCode = 500
		r.Error = result.Error
		r.Template = "error.tmpl"
		send(r, c)
		return
	}
	r.StatusCode = 200
	path := c.FullPath()
	if strings.HasSuffix(path, "/list") {
		path = path[:len("/list")]
	}
	r.Data = gin.H{"rm": rm, "path": path}
	r.Template = "results.tmpl"
	send(r, c)

}

// gt get one record by id C{R}UD
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

// up update record CR{U}D
// NOTE: destructive, ie unset fields will set to zero value
func up(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	p.ID = uint(uintID)
	result := DB.Save(&p)
	if result.Error != nil {
		panic(result.Error)
	}
	//TODO:use send rc
	c.Set("message", "updated")
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}
}

// up2 update record CR{U}D
// NOTE: constructive ie unset fields will be left unchanged
func up2(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	id := c.Param("id")
	DB.Model(&p).Where("ID=?", id).Updates(&p)
	//TODO:use send rc
	c.Set("message", "updated")
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}
}

// dl delete record CRU{D}
// TODO:make another soft delete handler and end point
func dl(c *gin.Context) {
	p := Proto()
	id := c.Param("id")
	result := DB.Delete(&p, id)
	fmt.Printf("%+v\n", result)
	c.Set("message", "deleted")
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}
}

// nw new record form
// TODO:return empty curl with json data to be filled
func nw(c *gin.Context) {
	p := Proto()
	r := &Reply{}
	formValues, err := StructFields(p, "form")
	if err != nil {
		r.StatusCode = 400
		r.Error = err
		r.Template = "error.tmpl"
		send(r, c)
		return
	}
	r.StatusCode = 200
	r.Data = formValues
	r.Template = "new.tmpl"
	send(r, c)
	return
}

// ed edit record form
// TODO:return empty curl with json data to be filled
func ed(c *gin.Context) {
	p := Proto()
	r := &Reply{}
	m := map[string]any{} //m record {m}ap
	//TODO:handle errorr
	id := c.Param("id")
	DB.Model(&p).First(&m, id)
	r.StatusCode = 200
	r.Data = gin.H{"m": m, "path": c.Request.URL.EscapedPath(), "id": id}
	r.Template = "edit.tmpl"
	send(r, c)
	return
}

// StructFields given struct and key
// returns fields tags values slice of that key
func StructFields(aStruct any, aKey string) ([]string, error) {
	values := []string{}
	typ := reflect.TypeOf(aStruct)
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s is not a struct", typ)
	}
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		if val := fld.Tag.Get(aKey); val != "" {
			values = append(values, val)
		}
	}
	return values, nil
}

// fragments buffer
//result := DB.First(&p, id)
//c.JSON(200, gin.H{"data": p})
//result := DB.Model(&p).Where("ID = ?", id).First(&m)
//c.JSON(200, gin.H{"data": p})
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
//wg.POST("", cr)
//wg.GET("", mw.KissAuth, rt)
//wg.GET("/:id", gt)
//wg.PUT("/:id", up)
//wg.POST("/:id/update", up)
//wg.POST("/:id/delete", dl)
//wg.DELETE("/:id", dl)
//wg.GET("/new", nw)
//wg.GET("/:id/edit", ed)
