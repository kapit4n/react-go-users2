package main

import (
	"net/http"
	"react-go-users2/controllers"
	"react-go-users2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	r.GET("/users", controllers.FindUsers)
	r.GET("/fizzbuzz/:n", controllers.FizzBuzzBuilder)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)

	r.Run()
}
