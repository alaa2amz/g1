package main

import (
	"fmt"
	//"os"
	"github.com/alaa2amz/g1/service"
	_ "github.com/alaa2amz/g1/service/post"
)

func init() {
	fmt.Println("main init")

}

func main() {
	fmt.Println("Hello, Web!")
	service.R.Run()
}
