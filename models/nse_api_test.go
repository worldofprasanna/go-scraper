package models_test

import (
	"testing"
	"fmt"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/worldofprasanna/go-scraper/models"
)
func TestFetchCommits(t *testing.T) {
	t.Run("should parse the response properly", func(t *testing.T) {
		nseScrapper := models.NewNSEScrapper("RBLBANK")
		boardMeeting, _ := nseScrapper.FetchBoardMeetingDetails()
		fmt.Println(boardMeeting)
		assert.Equal(t, "Fri Nov 29, 2019", boardMeeting.ParsedMeetingDate(), "should have the proper meeting date")
		assert.Equal(t, "Fund Raising", boardMeeting.Purpose, "should have the proper purpose")
		assert.Equal(t, "RBL Bank Limited", boardMeeting.CompanyName, "should have the proper company name")
		assert.Equal(t, "To consider Fund Raising Pursuant to Regulation 29 of the SEBI Listing Regulations, we would like to inform you that a meeting of the Board of Directors of the Bank is proposed to be held on Saturday, November 30, 2019, to inter alia consider and if thought fit to approve: Raising of funds by way of issue of equity shares of the Bank on a Preferential basis in accordance with the provisions of the Companies Act, 2013, Securities and Exchange Board of India (Issue of Capital and Disclosure Requirements) Regulations, 2018 and such other acts, rules and regulations as may be applicable, subject to the approval of the Members of the Bank.", boardMeeting.Details, "should have the proper details")
	})

	t.Run("should parse the meeting date properly", func(t *testing.T) {
		boardMeeting := models.BoardMeeting{
			MeetingDate: time.Now(),
		}
		assert.Equal(t, "Fri Nov 29, 2019", boardMeeting.ParsedMeetingDate(), "should parse the board meeting date properly")
	})
}
