package mw

import (
	//	"fmt"
	"github.com/alaa2amz/g1/helpers/ajwt"
	"github.com/alaa2amz/g1/service"
	"github.com/alaa2amz/g1/service/model"

	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func TrivH(c *gin.Context) {
	log.Println("b4")
	c.Next()
	log.Print("after")
}

func KissAuth(c *gin.Context) {
	//token := "123abc"
	authHeader := c.Request.Header["Authorization"]
	log.Println("zzzzz", authHeader)
	log.Println("b4")
	c.Next()
	log.Print("after")
}

func Triv() gin.HandlerFunc {
	return TrivH

}

//Logged only logged users
func Logged(c *gin.Context) {
	//authHeader := authHeader:=c.Request.Header["Authorization"]
	log.Println("logged-mw")
	authToken := c.GetHeader("Authorization")
	if authToken != "" {
		authToken = strings.Split(authToken, " ")[1]
	}

	coockiToken, coockiErr := c.Cookie("token")
	if coockiErr != nil {
		log.Print(coockiErr)
	}

	tokenStruct, err := ajwt.Valid(authToken)
	if err != nil {
		tokenStruct, err = ajwt.Valid(coockiToken)
	}
	if err != nil {
		if api := strings.Contains(c.Request.URL.Path, "/api/"); api {
			c.JSON(400, gin.H{"error": "not authorized"})
			c.Abort()
			return
		}
		referer := c.Request.URL.EscapedPath()
		c.Redirect(303, "login/new?re="+referer)

		return
	}
	c.Set("jwt", tokenStruct)
	userName, err := tokenStruct.Claims.GetSubject()
	if err != nil {
		c.AbortWithError(400, err)
	}
	User := model.User{}
	result := service.DB.Where("Name = ?", userName).First(&User)
	if result.Error != nil {
		c.AbortWithError(400, result.Error)
	}

	c.Set("UserID", User.ID)
	c.Set("User", User)

	c.Next()
}
