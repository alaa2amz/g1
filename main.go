// G1: Golang Generic Back End Rapid Builder.
// G1: Golang Generic Back End Rapid Builder.
// G1: Golang Generic Back End Rapid Builder.
// using Gin-Gorm as a wrapper.
// using gin and gorm.
// currently working on developing,
// framework like,
// backend rapid builder.
// Inspired by the philosophy of REACT.
// to make resources component like,
// i.e modular,
// out of the box,
// full crud restful API,
// web backend admin panel,
// for any resouce created resource,
// generated automatically
// any resource can be instantiated from any current resource,
// and easily moved from one back end to another.
// features:
//   - gin
//   - gorm
//   - unit tests using testify
//   - playwright front end tests(TODO)
//   - restful api
//   - web admin panel
//   - using html/template to render administration panel
//   - adaptive automatically generated frontend panel for
//
// every added or modified database table or resource
// …. and many more
// sample backend
// https://alaazak.alwaysdata.net/
// please ask by mail to see source code
package main

import (
	"embed"
	"fmt"
	"html/template"

	"github.com/alaa2amz/g1/service"
	//"github.com/gin-contrib/static"
	//"github.com/gin-gonic/gin"

	_ "github.com/alaa2amz/g1/service/component/login"
	_ "github.com/alaa2amz/g1/service/component/post"
	_ "github.com/alaa2amz/g1/service/component/egg"
	_ "github.com/alaa2amz/g1/service/component/wig"
	_ "github.com/alaa2amz/g1/service/component/tag"
	_ "github.com/alaa2amz/g1/service/component/user"
	_ "github.com/alaa2amz/g1/service/component/comment"
)

//go:embed tmpl static
var content embed.FS

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("Hello, Web!", service.Index)

	tmpl, err := template.ParseFS(content, "tmpl/*.tmpl")
	if err != nil {
		panic(err)
	}

	service.R.SetHTMLTemplate(tmpl)
	
	service.PostMigrate()
	
	service.R.Run()
}



//=========================================
	/*
	m := map[string]any{}
	service.DB.Table("posts").Take(&m)
	fmt.Println(m)
	*/
	//embedding static issue
	/* 
	service.R.StaticFS("/static", http.FS(content))
	service.R.Use(static.Serve("/", static.EmbedFolder(content, "static/")))
	service.R.NoRoute(func(c *gin.Context) {
		fmt.Printf("%s doesn't exists, redirect on /\n", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	*/
	/*
	cols, _ := service.DB.Migrator().ColumnTypes("posts")
	for _, n := range cols {
		fmt.Println(n.Name())
	}
	*/
	//fmt.Printf("\033[31mChilds:%+v\n",service.Childs)

//dgdgdgdgdgdg
