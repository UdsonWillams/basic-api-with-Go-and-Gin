package router

import (
	"github.com/UdsonWillams/basic-api-with-Go-and-Gin/handler"
	"github.com/gin-gonic/gin"
)

func home(router *gin.Engine) {

	v1_albums := router.Group("/api/v1/albums")
	v1_words := router.Group("/api/v1/words")

	{
		v1_albums.GET("/", handler.GetAlbums)
		v1_albums.GET("/:id", handler.GetAlbumByID)
		v1_albums.POST("/", handler.PostAlbums)
	}

	// My own examples.
	{
		v1_words.POST("/order", handler.PostOrderWords)
		v1_words.POST("/count-vowels", handler.PostCountVowels)
	}

}
