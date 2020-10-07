package controllers

import (
	"net/http"
	"react-go-users2/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name string `json: "name", binding:""required`
}

func FindUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FixzBuzzBuilder(c *gin.Context) {
	n, _ := strconv.Atoi(c.Param("n"))
	res := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 && i%7 == 0 {
			res = append(res, "FizzBuzzJazz")
		} else if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 && i%7 == 0 {
			res = append(res, "FizzJazz")
		} else if i%5 == 0 && i%7 == 0 {
			res = append(res, "BuzzJazz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else if i%7 == 0 {
			res = append(res, "Jazz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": res})

}

func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
