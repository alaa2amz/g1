// Alaa Ahmed Zakariya.
// Post ressource.
// Example component.
// a Social text.
// holds model, web , api ,test.
// TODO:generic joiner
// TODO:edit and new points
// TODO:validate basic data types on create
// TODO:intensive error handling and testing
package login

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/alaa2amz/g1/mw"
	"github.com/alaa2amz/g1/service"
	"github.com/alaa2amz/g1/service/tag"

	//	"github.com/mitchellh/mapstructure"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/golang-jwt/jwt/v5"
)

var (
	DB             *gorm.DB
	Path           string = "/g1/login"
	Editions              = []string{"", "/api"}
	TemplatePath          = "login/template"
	//TODO:move near model
	DroppedColumns  []string
	 secretKey = []byte("your-secret-key")
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

type Login struct {
	ID    uint   `form:"id" gorm:"primaryKey"` //id should be removed from form
	Username string `form:"username" binding:"required"`
	Password string `form:"password" gorm:"default:null;not null"`
	Token string `gorm:"default:null;not null"`
	ExpiredAt     *time.Time `time_format:"2006-01-02"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt   time.Time
	//gorm.Model

}
type Tag tag.Tag

func Proto() (p Login) { return }

func Protos() (p []Login) { return }

func init() {
	log.Println(Path + "init")
	if service.DB == nil {
		log.Fatal("main database not initialized")
	}
	DB = service.DB
	for _, dropColumn := range DroppedColumns {
		DB.Migrator().DropColumn(Proto(), dropColumn)
	}
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
	//r.LoadHTMLGlob("**/login/**/*.tmpl")

	for _, edition := range Editions {
		fullPath := edition + Path
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
		r.GET(fullPath, mw.KissAuth, rt)
		r.GET(fullPath+"/:id", gt)
		r.PUT(fullPath+"/:id", up)
		r.POST(fullPath+"/:id/update", up)
		r.POST(fullPath+"/:id/delete", dl)
		r.DELETE(fullPath+"/:id", dl)

		//r.GET(fullPath+"/test", tst)
		r.GET(fullPath+"/new", nw)
		r.GET(fullPath+"/:id/edit", ed)
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
	/////////////////////////////////////////////////////////////
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "alaa",                    // Subject (user identifier)
		"iss": "alaazak",                  // Issuer
		"aud": "client",           // Audience (user role)
		"exp": time.Now().Add(time.Second*60*5).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})
	fmt.Printf("Token claims added: %+v\n", token)
	tokenString, errt := token.SignedString(secretKey)
    	if errt != nil {
        	panic( errt)
    	}
	fmt.Printf("Token claims added: %+v\n", tokenString)
    	//////////////////////////////////////////////////////////////////////
	p.Token = tokenString

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
	c.Set("message", "logged in")
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, "/g1/post")
		return
	}

}

// rt retrieve records C{R}UD
func rt(c *gin.Context) {
	//TODO:sort columns and rows
	//TODO:smart sort
	//TODO:pagination
	ps := Protos() 		//model {p}rototype{s}
	rm := []map[string]any{} //{r}esults {m}ap
	r := &Reply{} 		//response reply
	qs := c.Request.URL.Query() //query string map
	terms := parseQueryString(qs)
	log.Printf("terms: %+v\n", terms) //debug
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
	r.Data = gin.H{"rm": rm, "path": c.FullPath()}
	r.Template = "results.tmpl"
	log.Printf("%+v\n",rm)
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
	DB.Delete(&p, id)
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
	r.Template = "login.tmpl"
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
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
	//authHeader:=c.Request.Header.Get("Authorization")
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Token missing in cookie")
		c.Redirect(303, "/g1/login/new")
		c.Abort()
		return
	}
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
// Check for verification errors
	if (err != nil)||(!token.Valid) {
		c.Redirect(303,"/g1/login/new")
		return
	}
	c.Next()
}}
// fragments buffer
//result := DB.First(&p, id)
//c.JSON(200, gin.H{"data": p})
//result := DB.Model(&p).Where("ID = ?", id).First(&m)
//c.JSON(200, gin.H{"data": p})
