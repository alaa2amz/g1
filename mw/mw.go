package mw

import "github.com/gin-gonic/gin"
import "log"

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
