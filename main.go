package main

import (
	"fmt"

	"github.com/alaa2amz/g1/service"
	_ "github.com/alaa2amz/g1/service/post"
	_ "github.com/alaa2amz/g1/service/login"
	_ "github.com/alaa2amz/g1/service/user"
	// "github.com/alaa2amz/g1/service/tag"
)

func init() {
	fmt.Println("main init")

}

func main() {
	fmt.Println("Hello, Web!")
	service.R.Run()
}
