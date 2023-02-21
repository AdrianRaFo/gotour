package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Email     string `json:"email"`
}

var users []User

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	v1.GET("/hello/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + c.Param("name") + "!",
		})
	})

	v1.POST("/signup", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			log.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
			return
		} else {
			users = append(users, newUser)
			c.JSON(http.StatusOK, users)
		}
	})

	router.Run(":8080")
}