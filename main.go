package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/worldofprasanna/go-scraper/models"
	log "github.com/sirupsen/logrus"
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
		scrapper := models.NewNSEScrapper(symbol)
		meetingInfo, err := scrapper.FetchBoardMeetingDetails()
		if err != nil {
			log.Error("error occurred while fetching data", err.Error())
			c.HTML(http.StatusOK, "result.tmpl", gin.H{
				"error": "Hmm ... Something wrong with machine to machine communication. Humans are looking into it.",
			})
		}
		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"symbol": symbol,
			"company_name": meetingInfo.CompanyName,
			"purpose": meetingInfo.Purpose,
			"details": meetingInfo.Details,
			"meeting_date": meetingInfo.ParsedMeetingDate(),
		})
	})
	r.Run()
}