package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", getHome)
	router.GET("/article/:title", getArticle)

	router.Run()
}

func getHome(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": "success",
		"message": "welcome gin",
	})
}

func getArticle(context *gin.Context) {
	title := context.Param("title")
	context.JSON(200, gin.H {
		"status": "success",
		"message": title,
	})
}
