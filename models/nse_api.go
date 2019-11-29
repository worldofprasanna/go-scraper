package models

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"regexp"
)

// NSEScrapper - Scrapper which can scrape the NSE website and fetch the information
type NSEScrapper struct {
	URL string
	Symbol string
}

// BoardMeeting - Struct used to capture the NSE Board Meeting values
type BoardMeeting struct {
	MeetingDate time.Time `json:"BoardMeetingDate"`
	Purpose string `json:"Purpose"`
	Details string `json:"Details"`
	CompanyName string `json:"CompanyName"`
}

var baseURL = "https://www.nseindia.com/corporates/corpInfo/equities/getBoardMeetings.jsp?Symbol=%s&Industry=&Period=Latest%%20Announced&Purpose=&period=Latest%%20Announced&symbol=%s&industry=&purpose="

// NewNSEScrapper - Creates new scrapper with the given symbol
func NewNSEScrapper(symbol string) *NSEScrapper {
	url := fmt.Sprintf(baseURL, symbol, symbol)
	return &NSEScrapper {
		URL: url,
		Symbol: symbol,
	}
}

// FetchBoardMeetingDetails - Fetches information about the board meeting
func (scrapper NSEScrapper) FetchBoardMeetingDetails() (*BoardMeeting, error) {
	resp, err := http.Get(scrapper.URL)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	boardMeetingInfo := BoardMeeting{}
	if err := boardMeetingInfo.UnmarshalJSON([]byte(body)); err != nil {
		log.Println(err)
		return nil, err
	}

	return &boardMeetingInfo, nil
}

// UnmarshalJSON - Override the UnmarshalJSON and parse the values accordingly
func (val *BoardMeeting) UnmarshalJSON(data []byte) error {

	body := string(data)
	layoutForDate := "2-Jan-2006"
	var err error
	companyNameRE := regexp.MustCompile("CompanyName:\"(.*?)\"")
	meetingDateRE := regexp.MustCompile("BoardMeetingDate:\"(.*?)\"")
	purposeRE := regexp.MustCompile("Purpose:\"(.*?)\"")
	detailsRE := regexp.MustCompile("Details:\"(.*?)\"")

	val.CompanyName = companyNameRE.FindStringSubmatch(body)[1]
	val.Purpose = purposeRE.FindStringSubmatch(body)[1]
	val.Details = detailsRE.FindStringSubmatch(body)[1]
	val.MeetingDate, err = time.Parse(layoutForDate, meetingDateRE.FindStringSubmatch(body)[1])

	return err
}

// ParsedMeetingDate - Fetches the meeting date in required format
func (val BoardMeeting) ParsedMeetingDate() string {
	return val.MeetingDate.Format("Mon Jan _2, 2006")
}