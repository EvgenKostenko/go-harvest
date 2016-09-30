package harvest

import (
	"fmt"
	"github.com/EvgenKostenko/go-harvest/models"
	"net/http"
)

type UserService struct {
	client *Client
}

//Type for users list
type People []struct {
	User models.User `json:"user"`
}

type UserDetail struct {
	User models.User `json:"user"`
}

type PeopleOptions struct {
	// You can also filter by updated_since to only show people that have been updated since the date you pass
	// UpdatedSince=2015-04-25+18%3A30
	UpdatedSince string `url:"updated_since,omitempty"`
}

func (s *UserService) People(opt *PeopleOptions) (*People, *http.Response, error) {
	apiEndpoint := "people"
	url, err := addOptions(apiEndpoint, opt)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	people := new(People)
	resp, err := s.client.Do(req, people)
	if err != nil {
		return nil, resp, err
	}

	return people, resp, err
}

func (s *UserService) GetUser(userId int) (*models.User, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("%s/%d", "people", userId)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(UserDetail)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return &user.User, resp, err
}
