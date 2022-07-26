package main

import (
	"gin/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Main website",
		"content": "test123",
	})
}

func main() {

	// 不需取值
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/", test)

	server.POST("/post", func(c *gin.Context) {

		//建立一個變數
		var json map[string]interface{}

		//將json與變數綁定
		c.BindJSON(&json)
		// fmt.Println(json)
		err := models.InsertUser(json)

		if err != nil {
			log.Fatal(err)
		}
	})

	server.Run(":8888")
}
