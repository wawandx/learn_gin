package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		articles := v1.Group("/article")
		{
			articles.GET("/", getHome)
			articles.GET("/:title", getArticle)
			articles.POST("/", postArticle)
		}
	}

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
