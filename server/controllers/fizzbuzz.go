package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FizzBuzzBuilder func
func FizzBuzzBuilder(c *gin.Context) {
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
