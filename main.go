package main

import (
	"fmt"
	"github.com/alaa2amz/g1/service"
	_ "github.com/alaa2amz/g1/service/component/comment"
	_ "github.com/alaa2amz/g1/service/component/co"
	_ "github.com/alaa2amz/g1/service/component/tag"
	_ "github.com/alaa2amz/g1/service/component/post"
	_ "github.com/alaa2amz/g1/service/component/user"
	_ "github.com/alaa2amz/g1/service/component/login"
)

func init() {
	fmt.Println("main init")

}

func main() {
	service.PostMigrate()
	fmt.Println("Hello, Web!",service.Index)
	m:=map[string]any{}
	service.DB.Table("posts").Take(&m)
	fmt.Println(m)
	cols,_:=service.DB.Migrator().ColumnTypes("posts")
	for _,n:=range cols{
		fmt.Println(n.Name())
	}
	service.R.Run()
}
