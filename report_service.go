package harvest

import (
	"fmt"
	"github.com/EvgenKostenko/go-harvest/models"
	"net/http"
)

type ReportService struct {
	client *Client
}

type DayEntryDetail struct {
	DayEntry models.DayEntry `json:"day_entry"`
}

type ReportOptions struct {
	// Requests can be filtered by From and To parameters
	// From=2015-04-25
	// To=2015-04-25
	// Acceptable values for the OnlyBilled, OnlyUnbilled, IsClosed  parameter are yes and no.
	// UpdatedSince=2015-04-25+18%3A30
	From         string `url:"from"`
	To           string `url:"to"`
	OnlyBilled   string `url:"only_billed,omitempty"`
	OnlyUnbilled string `url:"only_unbilled,omitempty"`
	IsClosed     string `url:"is_closed,omitempty"`
	UpdatedSince string `url:"updated_since,omitempty"`
	UserId       string `url:"user_id,omitempty"`
}

//Type for day entry list
type DayEntries []struct {
	DayEntry models.DayEntry `json:"day_entry"`
}

func (s *ReportService) DayEntriesByProject(projectId int, opt *ReportOptions) (*DayEntries, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("projects/%d/entries", projectId)

	dayEntries, resp, err := s.requestDayEntries(apiEndpoint, opt)

	return dayEntries, resp, err
}

func (s *ReportService) DayEntriesByUser(userId int, opt *ReportOptions) (*DayEntries, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("people/%d/entries", userId)

	dayEntries, resp, err := s.requestDayEntries(apiEndpoint, opt)

	return dayEntries, resp, err
}

func (s *ReportService) requestDayEntries(apiEndpoint string, opt *ReportOptions) (*DayEntries, *http.Response, error) {
	url, err := addOptions(apiEndpoint, opt)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	dayEntries := new(DayEntries)
	resp, err := s.client.Do(req, dayEntries)

	if err != nil {
		return nil, resp, err
	}

	return dayEntries, resp, err
}