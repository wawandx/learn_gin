package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H {
			"message": "welcome gin",
		})
	})

	route.Run()
}