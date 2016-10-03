package harvest

import (
	"github.com/EvgenKostenko/go-harvest/models"
	"net/http"
)

type ProjectService struct {
	client *Client
}

type ProjectOptions struct {
	// Requests can be filtered by client_id and updated_since
	// UpdatedSince=2015-04-25+18%3A30
	Client string `url:"client,omitempty"`
	UpdatedSince string `url:"updated_since,omitempty"`
}

//Type for users list
type Projects []struct {
	Project models.Project `json:"project"`
}

func (s *ProjectService) Projects(opt *ProjectOptions) (*Projects, *http.Response, error) {
	apiEndpoint := "projects"
	url, err := addOptions(apiEndpoint, opt)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	projects := new(Projects)
	resp, err := s.client.Do(req, projects)
	if err != nil {
		return nil, resp, err
	}

	return projects, resp, err
}