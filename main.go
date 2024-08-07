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
	service.R.Run()
}
