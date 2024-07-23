package mw

import (
	"github.com/gin-gonic/gin"
 "log"
//import "github.com/alaa2amz/ajwt"
 "strings"
 "fmt"
 )

func TrivH(c *gin.Context) {
	log.Println("b4")
	c.Next()
	log.Print("after")
}


func KissAuth(c *gin.Context) {
	//token := "123abc"
	authHeader:=c.Request.Header["Authorization"]
log.Println("zzzzz",authHeader)
	log.Println("b4")
	c.Next()
	log.Print("after")
}



func Triv() gin.HandlerFunc {
	return TrivH

}

func Logged(c *gin.Context) {
	//authHeader := authHeader:=c.Request.Header["Authorization"]
	authToken:=c.GetHeader("Authorization")
	if authToken != "" {
		authToken = strings.Split(authToken," ")[1]
	}
	log.Print(authToken)
	fmt.Println(authToken)
	coockiToken ,err:= c.Cookie("token")
	if err != nil {
		log.Print(err)
	}
	log.Print(coockiToken)
	c.Next()
}
