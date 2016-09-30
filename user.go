package harvest

import (
	"time"
	"net/http"
)


type UserService struct {
	client *Client
}

//Type for users list
type People []struct {
	User User `json:"user"`
}

type User struct {
	ID                           int       `json:"id"`
	Email                        string    `json:"email"`
	CreatedAt                    time.Time `json:"created_at"`
	IsAdmin                      bool      `json:"is_admin"`
	FirstName                    string    `json:"first_name"`
	LastName                     string    `json:"last_name"`
	Timezone                     string    `json:"timezone"`
	IsContractor                 bool      `json:"is_contractor"`
	Telephone                    string    `json:"telephone"`
	IsActive                     bool      `json:"is_active"`
	HasAccessToAllFutureProjects bool      `json:"has_access_to_all_future_projects"`
	DefaultHourlyRate            int       `json:"default_hourly_rate"`
	Department                   string    `json:"department"`
	WantsNewsletter              bool      `json:"wants_newsletter"`
	UpdatedAt                    time.Time `json:"updated_at"`
	CostRate                     string    `json:"cost_rate"`
	IdentityAccountID            int       `json:"identity_account_id"`
	IdentityUserID               int       `json:"identity_user_id"`
}



func (s *UserService) People() (*People, *http.Response, error) {
	apiEndpoint := "people"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
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