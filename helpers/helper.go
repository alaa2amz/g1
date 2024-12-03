package helpers

import (
	"fmt"
	"log"
	"reflect"
	"slices"
	"strings"

	"github.com/alaa2amz/g1/helpers/ajwt"
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Reply struct {
	StatusCode int
	Data       any
	Error      error
	Template   string
	Redirect   string
}

type QueryTerm struct {
	Or       bool
	Column   string
	Relation string
	//Values   string
	Values []string
}

var Relations = map[string]string{
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
func Send(r *Reply, c *gin.Context) {
	errMsg := ""
	if r.Error != nil {
		errMsg = r.Error.Error()
	}
	if ok := strings.Contains(c.Request.URL.Path, "/api/"); ok {
		//handeling data and errors
		c.JSON(r.StatusCode, gin.H{"data": r.Data, "error": errMsg})
		return
	} else if c.Request.Method != "GET" &&
		r.Error == nil &&
		r.Redirect != "" &&
		!strings.Contains(c.Request.URL.Path, "/api/") &&
		r.StatusCode >= 200 && r.StatusCode < 300 {
		c.Redirect(303, r.Redirect)
	} else {

		//TODO:handle templates errors and unfound templates or empty templates
		c.HTML(r.StatusCode, r.Template, gin.H{"data": r.Data, "error": errMsg})
		return
	}
}

func ParseQueryString(qs map[string][]string) []QueryTerm {
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
			if v, ok := Relations[vals[0]]; ok {
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

func TidySlice(orig []string, leads, trails []string) []string {
	var middleKeys, bkbKeys, keys []string
	for _, key := range leads {
		if ok := slices.Contains(orig, key); !ok {
			i := slices.Index(leads, key)
			if i < 0 {
				panic("strange error")
			}
			leads = slices.Delete(leads, i, i+1)
		}
	}

	for _, key := range trails {
		if ok := slices.Contains(orig, key); !ok {
			i := slices.Index(trails, key)
			if i < 0 {
				panic("strange error")
			}
			trails = slices.Delete(trails, i, i+1)
		}
	}
	for _, key := range orig {
		bkbKeys = append(bkbKeys, key)
		if (!slices.Contains(leads, key)) && (!slices.Contains(trails, key)) {
			middleKeys = append(middleKeys, key)
		}
	}
	if len(leads)+len(middleKeys)+len(trails) < len(orig) {
		log.Panicln("col ordering error")
		return bkbKeys
	}
	keys = append(keys, leads...)
	keys = append(keys, middleKeys...)
	keys = append(keys, trails...)
	return keys
}

func Cr(c *gin.Context, db *gorm.DB, f func() any) {
	r := &Reply{}
	p := f()

	err := c.ShouldBind(&p)
	if err != nil {
		r.StatusCode = 400
		r.Error = err
		r.Template = "error.tmpl"
		Send(r, c)
		return
	}

	result := db.Create(&p)
	if result.Error != nil {
		r.StatusCode = 400
		r.Error = result.Error
		r.Template = "error.tmpl"
		Send(r, c)
		return
	}
	//

	//TODO:use send rc
	c.Set("message", "updated")
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, c.Request.URL.Path)
		return
	}

}
func Bail(c *gin.Context, code int, err error) {
	if err != nil {
		r := &Reply{}
		r.StatusCode = code
		r.Error = err
		r.Template = "error.tmpl"
		Send(r, c)
		c.AbortWithError(code, err)
		return
	}
}

func UserCrPass(p *model.User) error {
	// /---///
	if p.Password != p.Confirm {

		return fmt.Errorf("Password not confirming")
	}
	var err error
	p.PH, err = bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	return nil
}

func LoginCrToken(p *model.Login, c *gin.Context, db *gorm.DB) (string, error) {
	logUser := model.User{}
	result := db.Where("name = ?", p.Name).First(&logUser)
	if result.Error != nil {
		return "", result.Error
	}
	p.User = logUser
	err := bcrypt.CompareHashAndPassword(logUser.PH, []byte(p.Password))
	if err != nil {
		return "", err
	}
	token, err := ajwt.Token(ajwt.EasyClaims(p.Name, "client", 360))
	//token,err:=ajwt.Token(ajwt.EasyClaims(p.Name+" "+fmt.Sprint(p.ID),"client",360))
	if err != nil {
		return "", err
	}
	//p.TH,err=bcrypt.GenerateFromPassword([]byte(token),bcrypt.MinCost)
	//if err != nil {
	//	return "",err
	//}
	c.SetCookie("token", token, 3600, "/", c.Request.Host, false, true)
	//c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.Set("message", "logged in")
	//c.Set("UserID",logUser.ID)
	//c.Set("User",logUser)
	return token, nil

}
