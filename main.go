package main

import (
	"log"
	"gin-user-service/models"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("====== gin user service launch ======")
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person models.Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Email)
		log.Println(person.Secret)
		log.Println(person.BirthDate)
	}
	c.String(200, "Success")
}
