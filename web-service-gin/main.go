package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordsBody struct {
	Words []string `json:"words"`
	Order string   `json:"order"`
}

type VowelsBody struct {
	Vowels []string `json:"vowels"`
}

type WordsOrder struct {
	asc  string
	desc string
}

func postOrderWords(c *gin.Context) {
	var orderWords WordsBody
	if err := c.BindJSON(&orderWords); err != nil {
		return
	}
	response := SortWords(orderWords.Words, orderWords.Order)
	c.IndentedJSON(http.StatusOK, response)
}

func postCountVowels(c *gin.Context) {
	var countVowels VowelsBody
	if err := c.BindJSON(&countVowels); err != nil {
		return
	}
	response := CountVowels(countVowels.Vowels)
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	// In Gin, the colon preceding an item in the path
	// signifies that the item is a path parameter.
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// My own examples.
	router.POST("/words/order", postOrderWords)
	router.POST("/words/count-vowels", postCountVowels)

	router.Run("localhost:8081")
}
