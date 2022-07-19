package main

import (
	"net/http"
	"regexp"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func Greetings(c *gin.Context){
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := c.Query("email")
	if Re.MatchString(email){
	c.IndentedJSON(http.StatusOK,"Email Verified")
	}else {
		c.IndentedJSON(http.StatusOK,"Email Not Verified")
	}
}

func main(){
	router:=gin.Default()
	router.Use(cors.Default())
	router.POST("/",Greetings)
	router.Run("localhost:8080")
}