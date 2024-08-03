package main

import (
	"fmt"
	//"os"
	"github.com/alaa2amz/g1/service"
	_ "github.com/alaa2amz/g1/service/comment"
	_ "github.com/alaa2amz/g1/service/tag"

	// _ "github.com/alaa2amz/g1/service/pay_method"
	_ "github.com/alaa2amz/g1/service/post"
	_ "github.com/alaa2amz/g1/service/user"
	_ "github.com/alaa2amz/g1/service/login"
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
