package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", getHome)
	router.GET("/article/:title", getArticle)
	router.POST("/articles", postArticle)

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

func postArticle(context *gin.Context) {
	title := context.PostForm("title")
	desc := context.PostForm("desc")

	context.JSON(200, gin.H {
		"status": "success",
		"title": title,
		"desc": desc,
	})
}
