package harvest

import (
	"fmt"
	"github.com/EvgenKostenko/go-harvest/models"
	"net/http"
)

type ProjectService struct {
	client *Client
}

type ProjectDetail struct {
	Project models.Project `json:"project"`
}

type ProjectOptions struct {
	// Requests can be filtered by client_id and updated_since
	// UpdatedSince=2015-04-25+18%3A30
	Client       string `url:"client,omitempty"`
	UpdatedSince string `url:"updated_since,omitempty"`
}

//Type for project list
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

func (s *ProjectService) GetProject(projectId int) (*models.Project, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("projects/%d", projectId)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(ProjectDetail)
	resp, err := s.client.Do(req, project)
	if err != nil {
		return nil, resp, err
	}

	return &project.Project, resp, err
}

//The method by which the project is budgeted or estimated.
// Parameter: BudgetBy or EstimateBy
// Options: Project (Hours Per Project),
//          Project_Cost (Total Project Fees),
// 			Task (Hours Per Task),
// 			Person (Hours Per Person),
// 			None (No Budget).\
func (s *ProjectService) CreateProject(project *ProjectDetail) (*http.Response, error) {
	apiEndpoint := "projects"

	resp, err := s.requestProject("POST", apiEndpoint, project)

	return resp, err
}

//The method by which the project is budgeted or estimated.
// Parameter: BudgetBy or EstimateBy
// Options: Project (Hours Per Project),
//          Project_Cost (Total Project Fees),
// 			Task (Hours Per Task),
// 			Person (Hours Per Person),
// 			None (No Budget).\
func (s *ProjectService) UpdateProject(project *ProjectDetail) (*http.Response, error) {
	apiEndpoint := fmt.Sprintf("projects/%d", project.Project.ID)

	resp, err := s.requestProject("PUT", apiEndpoint, project)

	return resp, err
}

func (s *ProjectService) requestProject(method, urlStr string, project *ProjectDetail) (*http.Response, error) {

	req, err := s.client.NewRequest(method, urlStr, project)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *ProjectService) DeleteProject(projectId int) (*http.Response, error) {
	apiEndpoint := fmt.Sprintf("projects/%d", projectId)

	req, err := s.client.NewRequest("DELETE", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
