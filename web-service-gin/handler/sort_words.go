package handler

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type WordsBody struct {
	Words []string `json:"words"`
	Order string   `json:"order"`
}

type WordsOrder struct {
	asc  string
	desc string
}

var ordenation WordsOrder = WordsOrder{asc: "asc", desc: "desc"}

func SortWords(value []string, order string) []string {
	if order == ordenation.desc {
		return sortWordsDesc(value)
	}
	return sortWordsAsc(value)

}

func sortWordsAsc(value []string) []string {
	sort.Strings(value)
	return value
}

func sortWordsDesc(value []string) []string {
	sort.Slice(value, func(i, j int) bool {
		// retorno dá função anonima
		return value[i] > value[j]
	})
	return value
}

func PostOrderWords(c *gin.Context) {
	var orderWords WordsBody
	if err := c.BindJSON(&orderWords); err != nil {
		return
	}
	response := SortWords(orderWords.Words, orderWords.Order)
	c.IndentedJSON(http.StatusOK, response)
}
