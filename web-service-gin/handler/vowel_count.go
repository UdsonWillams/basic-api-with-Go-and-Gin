package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VowelsBody struct {
	Vowels []string `json:"vowels"`
}

func countVowels(word []string) map[string]map[string]int {
	total_slice := map[string]int{}
	for _, value := range word {
		total := 0
		for _, v := range value {
			switch string(v) {
			case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
				total++
			}
		}
		total_slice[value] = total
	}

	response := map[string]map[string]int{"total_vowels": total_slice}
	return response
}

func PostCountVowels(c *gin.Context) {
	var count_vowels VowelsBody
	if err := c.BindJSON(&count_vowels); err != nil {
		return
	}
	response := countVowels(count_vowels.Vowels)
	c.IndentedJSON(http.StatusOK, response)
}
