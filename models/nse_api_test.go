package models_test

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/worldofprasanna/go-scraper/models"
	"net/http/httptest"
	"net/http"
)
func TestNSEApi(t *testing.T) {
	t.Run("should parse the response properly", func(t *testing.T) {
		mockResponse := `{ success:true ,results:1,rows:[{Symbol:"RBLBANK",CompanyName:"RBL Bank Limited",ISIN:"INE976G01028",Ind:"-",Purpose:"Fund Raising",BoardMeetingDate:"30-Nov-2019",DisplayDate:"27-Nov-2019",seqId:"103950873",Details:"Some Details"}]}`
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(mockResponse))
		}))
		defer server.Close()

		nseScrapper := models.NSEScrapper{
			URL: server.URL,
			Client: server.Client(),
			Symbol: "AnySymbol",
		}
		boardMeeting, _ := nseScrapper.FetchBoardMeetingDetails()
		assert.Equal(t, "Sat Nov 30, 2019", boardMeeting.ParsedMeetingDate(), "should have the proper meeting date")
		assert.Equal(t, "Fund Raising", boardMeeting.Purpose, "should have the proper purpose")
		assert.Equal(t, "RBL Bank Limited", boardMeeting.CompanyName, "should have the proper company name")
		assert.Equal(t, "Some Details", boardMeeting.Details, "should have the proper details")
	})

	t.Run("should parse the meeting date properly", func(t *testing.T) {
		boardMeeting := models.BoardMeeting{
			MeetingDate: time.Date(2019, time.November, 30, 23, 0, 0, 0, time.UTC),
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
