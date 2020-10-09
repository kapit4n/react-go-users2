package main

import (
	"net/http"
	"react-go-users2/controllers"
	"react-go-users2/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	r.GET("/fizzbuzz/:n", controllers.FizzBuzzBuilder)

	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)

	r.GET("/news", controllers.FindNewsDefault)
	r.GET("/news/:search", controllers.FindNews)

	r.Run()
}
