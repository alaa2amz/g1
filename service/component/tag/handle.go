// Alaa Ahmed Zakariya.
// Post ressource.
// Example component.
// a Social text.
// holds model, web , api ,test.
// TODO:generic joiner
// TODO:edit and new points
// TODO:validate basic data types on create
// TODO:intensive error handling and testing
package tag

import (
	"fmt"
	h "github.com/alaa2amz/g1/helpers"
	"github.com/alaa2amz/g1/service/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	//"github.com/alaa2amz/g1/mw"
	//"golang.org/x/exp/maps"
	//"github.com/alaa2amz/g1/service/model"
	//"github.com/mitchellh/mapstructure"
	//"reflect"
)

func tst(c *gin.Context) {
	r := &h.Reply{}
	h.Send(r, c)
}

// cr Create Handler
func cr(c *gin.Context) {
	p := Proto()
	err := c.ShouldBind(&p)
	h.Bail(c, 400, err)

	///---
	//id := c.GetUint("UserID")
	//p.UserID = &id
	///---///

	result := DB.Create(&p)
	h.Bail(c, 400, result.Error)

	r := &h.Reply{}
	r.StatusCode = 200
	r.Data = gin.H{"data": p}
	r.Redirect = Path
	h.Send(r, c)

}

// rt retrieve records C{R}UD
func rt(c *gin.Context) {
	//TODO:sort columns and rows
	//TODO:smart sort
	//TODO:pagination
	ps := Protos()              //model {p}rototype{s}
	rm := []map[string]any{}    //{r}esults {m}ap
	r := &h.Reply{}             //response reply
	qs := c.Request.URL.Query() //query string map
	terms := h.ParseQueryString(qs)
	log.Printf("terms: %+v\n", terms)  //debug
	QDB := DB.Session(&gorm.Session{}) //start session Query DB
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
	h.Bail(c, 500, result.Error)

	r.StatusCode = 200
	path := c.FullPath() //TODO:Unifi Path or c.FullPath
	if strings.HasSuffix(path, "/list") {
		path = path[:len("/list")]
	}
	log.Println("path:", path)
	keys := []string{}
	if len(rm) > 0 {
		keys = TidyCols
		log.Println(keys)
	}
	r.Data = gin.H{"rm": rm, "keys": keys, "path": path}
	r.Template = "results.tmpl"
	h.Send(r, c)

}

// gt get one record by id C{R}UD
func gt(c *gin.Context) {
	p := Proto()
	m := map[string]interface{}{}
	id := c.Param("id")
	result := DB.Model(&p).First(&m, id)
	h.Bail(c, 500, result.Error)
	/*
		if result.Error != nil {
			//c.JSON(400, gin.H{"error": result.Error.Error()})
			r := &h.Reply{400, nil, result.Error, "error.tmpl"}
			h.Send(r, c)
			return
		}
	*/

	r := &h.Reply{200, &m, nil, "show.tmpl", ""}
	h.Send(r, c)

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
	h.Bail(c, 400, result.Error)
	c.Set("message", "Updated")
	r := &h.Reply{}
	r.StatusCode = 200
	r.Data = gin.H{"data": p}
	r.Redirect = Path
	h.Send(r, c)
}

// up2 update record CR{U}D
// NOTE: constructive ie unset fields will be left unchanged
func up2(c *gin.Context) {
	p := Proto()
	c.Bind(&p)
	id := c.Param("id")
	result := DB.Model(&p).Where("ID=?", id).Updates(&p)
	h.Bail(c, 400, result.Error)
	c.Set("message", "Updated")
	r := &h.Reply{}
	r.StatusCode = 200
	r.Data = gin.H{"data": p}
	r.Redirect = Path
	h.Send(r, c)
	//TODO:use send rc
}

// dl delete record CRU{D}
// TODO:make another soft delete handler and end point
func dl(c *gin.Context) {
	p := Proto()
	id := c.Param("id")
	result := DB.Delete(&p, id)
	h.Bail(c, 400, result.Error)
	c.Set("message", "deleted")
	r := &h.Reply{}
	r.StatusCode = 200
	r.Data = gin.H{"data": p}
	r.Redirect = Path
	h.Send(r, c)
	return
}

// nw new record form
// TODO:return empty curl with json data to be filled
func nw(c *gin.Context) {
	//p := Proto()
	r := &h.Reply{}
	r.StatusCode = 200
	r.Data = gin.H{"cols": TidyCols, "path": Path}
	r.Template = "new.tmpl"
	h.Send(r, c)
	return
}

// ed edit record form
// TODO:return empty curl with json data to be filled
func ed(c *gin.Context) {
	p := Proto()
	r := &h.Reply{}
	m := map[string]any{} //m record {m}ap
	id := c.Param("id")
	result := DB.Model(&p).First(&m, id)
	h.Bail(c, 400, result.Error)
	r.StatusCode = 200
	//TODO:unify pathes
	r.Data = gin.H{"m": m, "path": c.Request.URL.EscapedPath(), "id": id}
	r.Template = "edit.tmpl"
	h.Send(r, c)
	return
}

func crAs(c *gin.Context) {
	//TODO:check pid if exist?
	pid := c.Param("id")
	ass := c.Param("path")
	log.Println(pid, ass)
	p := Proto()
	pidint, err := strconv.Atoi(pid)
	if err != nil {
		panic(err)
	}
	p.ID = uint(pidint)
	switch ass {
	case "comment":
		col := "Comments"      ///
		asP := model.Comment{} ///
		err = c.ShouldBind(&asP)
		if err != nil {
			panic(err)
		}
		err = DB.Model(&p).Association(col).Append(&asP)
		if err != nil {
			panic(err)
		}
	}
}

//func crAs2(c *gin.Context){

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
//TODO:use send rc
/*
	c.Set("message", "created")
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}
*/
/*

	if result.Error != nil {
		r.StatusCode = 500
		r.Error = result.Error
		r.Template = "error.tmpl"
		h.Send(r, c)
		return
	}
*/
//keys = tidySlice(maps.Keys(rm[0]), LeadCols, TrailCols)
////formValues, err := StructFields(p, "form")
//if err != nil {
//	r.StatusCode = 400
//	r.Error = err
//	r.Template = "error.tmpl"
//	send(r, c)
//	return
//i}
/*
	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}
*/
/*

	if strings.Contains(c.Request.URL.Path, "/api/") {
		c.JSON(200, gin.H{"data": p})
		return
	} else {
		c.Redirect(303, Path)
		return
	}
*/
/*
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
	}*/
