package models_test

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/worldofprasanna/go-scraper/models"
)
func TestNSEApi(t *testing.T) {
	t.Run("should parse the response properly", func(t *testing.T) {
		t.Skip("Mock the API call and use it")
		nseScrapper := models.NewNSEScrapper("RBLBANK")
		boardMeeting, _ := nseScrapper.FetchBoardMeetingDetails()
		assert.Equal(t, "Fri Nov 29, 2019", boardMeeting.ParsedMeetingDate(), "should have the proper meeting date")
		assert.Equal(t, "Fund Raising", boardMeeting.Purpose, "should have the proper purpose")
		assert.Equal(t, "RBL Bank Limited", boardMeeting.CompanyName, "should have the proper company name")
		assert.Equal(t, "To consider Fund Raising Pursuant to Regulation 29 of the SEBI Listing Regulations, we would like to inform you that a meeting of the Board of Directors of the Bank is proposed to be held on Saturday, November 30, 2019, to inter alia consider and if thought fit to approve: Raising of funds by way of issue of equity shares of the Bank on a Preferential basis in accordance with the provisions of the Companies Act, 2013, Securities and Exchange Board of India (Issue of Capital and Disclosure Requirements) Regulations, 2018 and such other acts, rules and regulations as may be applicable, subject to the approval of the Members of the Bank.", boardMeeting.Details, "should have the proper details")
	})

	t.Run("should parse the meeting date properly", func(t *testing.T) {
		boardMeeting := models.BoardMeeting{
			MeetingDate: time.Now(),
		}
		assert.Equal(t, "Sat Nov 30, 2019", boardMeeting.ParsedMeetingDate(), "should parse the board meeting date properly")
	})

	t.Run("should marshall the data properly", func(t *testing.T) {
		body := `{ success:true ,results:1,rows:[{Symbol:"RBLBANK",CompanyName:"RBL Bank Limited",ISIN:"INE976G01028",Ind:"-",Purpose:"Fund Raising",BoardMeetingDate:"30-Nov-2019",DisplayDate:"27-Nov-2019",seqId:"103950873",Details:"Some Details"}]}`
		boardMeeting := new(models.BoardMeeting)
		err := boardMeeting.UnmarshalJSON([]byte(body))
		assert.Nil(t, err, "should not contain any error")
		assert.Equal(t, "Sat Nov 30, 2019", boardMeeting.ParsedMeetingDate(), "should parse the board meeting date properly")
		assert.Equal(t, "Some Details", boardMeeting.Details, "should parse the details properly")
		assert.Equal(t, "Fund Raising", boardMeeting.Purpose, "should parse the purpose properly")
		assert.Equal(t, "RBL Bank Limited", boardMeeting.CompanyName, "should parse the company name properly")
	})

	t.Run("should return error if the results is not a valid number", func(t *testing.T) {
		body := `{ success:true ,results:nil, rows:[]}`
		boardMeeting := new(models.BoardMeeting)
		err := boardMeeting.UnmarshalJSON([]byte(body))
		assert.NotNil(t, err, "should return error")
		assert.Contains(t, err.Error(), "invalid syntax", "should return error with proper message")
	})

	t.Run("should return no data found error if the results are zero", func(t *testing.T) {
		body := `{ success:true ,results:0, rows:[]}`
		boardMeeting := new(models.BoardMeeting)
		err := boardMeeting.UnmarshalJSON([]byte(body))
		assert.NotNil(t, err, "should return error")
		assert.Contains(t, err.Error(), "No meetings found for the company.", "should return error with proper message")
	})
}
