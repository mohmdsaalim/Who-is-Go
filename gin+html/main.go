package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"name": "Saalim",
		"role": "Go Developer",
	})
})
	r.Run(":8800")
}
