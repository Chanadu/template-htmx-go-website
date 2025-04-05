package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*.html")
	e.Static("/css", "./public/css")

	e.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"name": "test",
		// })
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "asdoaksdj",
		})
	})

	err := e.Run(":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
