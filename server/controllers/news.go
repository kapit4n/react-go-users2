package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindNewsDefault(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"res": "Res"})
}

type Article struct {
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishAt   string
	Content     string
}

// NewsInfo
type NewsInfo struct {
	Status       string
	TotalResults int
	Articles     []Article
}

func FindNews(c *gin.Context) {
	// search := c.Param("search")
	resp, err := http.Get("http://newsapi.org/v2/everything?q=bitcoin&from=2020-09-09&sortBy=publishedAt&apiKey=77e82ace6f6c4dfaadf3836b68f14b4c")
	if err != nil {
		fmt.Printf("The request has errors")
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		newsInfo := NewsInfo{}
		// textBytes := []byte(resp.Body)

		err := json.Unmarshal(data, &newsInfo)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, newsInfo)
	}
}
