package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to the NSE Scraper",
		})
	})
	r.POST("/board_meeting", func(c *gin.Context) {
		symbol := c.PostForm("symbol")
		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"title": symbol,
		})
	})
	r.Run()
}