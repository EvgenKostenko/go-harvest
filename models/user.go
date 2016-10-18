package models

import "time"

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
	DefaultHourlyRate            float64   `json:"default_hourly_rate"`
	Department                   string    `json:"department"`
	WantsNewsletter              bool      `json:"wants_newsletter"`
	UpdatedAt                    time.Time `json:"updated_at"`
	CostRate                     float64   `json:"cost_rate"`
	WeeklyCapacity               int       `json:"weekly_capacity,omitempty"`
	SignupRedirectionCookie      string    `json:"signup_redirection_cookie,omitempty"`
}

type UserParameters struct {
	ID                           int     `json:"id,omitempty"`
	Email                        string  `json:"email,omitempty"`
	IsAdmin                      bool    `json:"is_admin,omitempty"`
	FirstName                    string  `json:"first_name,omitempty"`
	LastName                     string  `json:"last_name,omitempty"`
	Timezone                     string  `json:"timezone,omitempty"`
	IsContractor                 bool    `json:"is_contractor,omitempty"`
	Telephone                    string  `json:"telephone,omitempty"`
	IsActive                     bool    `json:"is_active,omitempty"`
	HasAccessToAllFutureProjects bool    `json:"has_access_to_all_future_projects,omitempty"`
	DefaultHourlyRate            float64 `json:"default_hourly_rate,omitempty"`
	Department                   string  `json:"department,omitempty"`
	CostRate                     float64 `json:"cost_rate,omitempty"`
}
