package main

import (
	"fmt"
	"github.com/alaa2amz/g1/service"
	_ "github.com/alaa2amz/g1/service/component/comment"
	_ "github.com/alaa2amz/g1/service/component/tag"
	_ "github.com/alaa2amz/g1/service/component/post"
	_ "github.com/alaa2amz/g1/service/component/user"
	_ "github.com/alaa2amz/g1/service/component/login"
)

func init() {
	fmt.Println("main init")

}

func main() {
	fmt.Println("Hello, Web!")
	for k, p := range service.Paths {
		fmt.Printf("%s -- %+v -- %T\n", k, p, p)
	}
	service.R.Run()
}
