package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/worldofprasanna/go-scraper/errorhandlers"
	"github.com/worldofprasanna/go-scraper/models"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": title,
		})
	})

	r.POST("/board_meeting", func(c *gin.Context) {
		symbol := c.PostForm("symbol")
		scrapper := models.NewNSEScrapper(symbol)
		meetingInfo, err := scrapper.FetchBoardMeetingDetails()
		if err != nil {
			log.Error("error occurred while fetching data: ", err.Error())
			userFriendlyMsg := "Hmm ... Something wrong with machine to machine communication. Humans are looking into it."
			noDataFoundErr, ok := err.(*errorhandlers.NoDataFound)
			if ok {
				userFriendlyMsg = noDataFoundErr.UserFriendlyMsg()
			}
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"error": userFriendlyMsg,
				"title": title,
			})
		} else {
			c.HTML(http.StatusOK, "result.tmpl", gin.H{
				"symbol":       symbol,
				"company_name": meetingInfo.CompanyName,
				"purpose":      meetingInfo.Purpose,
				"details":      meetingInfo.Details,
				"meeting_date": meetingInfo.ParsedMeetingDate(),
			})
		}
	})
	r.Run()
}

var title = "Welcome to the NSE Scraper"
