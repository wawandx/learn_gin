package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
 
type Article struct {
	gorm.Model
	Title string
	Slug string `gorm: "unique_index"`
	Desc string `sql: "type:text;"`
}

var DB *gorm.DB

func main() {
	var err error

	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:8889)/learn-gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	// Migrate the schema
  DB.AutoMigrate(&Article{})

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
	items := []Article{}
	DB.Find(&items)

	context.JSON(200, gin.H {
		"status": "success",
		"data": items,
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
