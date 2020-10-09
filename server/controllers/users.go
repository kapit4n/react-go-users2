package controllers

import (
	"net/http"
	"react-go-users2/models"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Name   string `json: "name"`
	Points int    `json: "points"`
}

func FindUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
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
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Points: input.Points}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	user.Name = input.Name
	user.Points = input.Points

	models.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
