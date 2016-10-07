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

type UserParameters struct {
	User models.UserParameters `json:"user"`
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

// Create new user
// POST https://YOURACCOUNT.harvestapp.com/people
// HTTP Response: 201 Created
// Upon creation, we’ll send an email to the user with instructions for setting a password.
// At minimum, you’ll need to include values for email, first-name, and last-name
func (s *UserService) CreateUser(user *models.UserParameters) (*http.Response, error) {
	apiEndpoint := "people"
	userDetail := UserParameters{User: *user}
	resp, err := s.requestUser("POST", apiEndpoint, &userDetail)

	return resp, err
}

// Update user by ID
// PUT https://YOURACCOUNT.harvestapp.com/people/{USERID}
// HTTP Response: 200 OK
// You can update selected attributes for a user with this request.
// Note, updates to password are disregarded.
func (s *UserService) UpdateUser(user *models.UserParameters) (*http.Response, error) {
	apiEndpoint := fmt.Sprintf("people/%d", user.ID)
	userDetail := UserParameters{User: *user}
	resp, err := s.requestUser("PUT", apiEndpoint, &userDetail)

	return resp, err
}

// This is universal method for create or update user
func (s *UserService) requestUser(method, urlStr string, user *UserParameters) (*http.Response, error) {

	req, err := s.client.NewRequest(method, urlStr, user)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
