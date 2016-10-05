package models

import "time"

type Project struct {
	ID                               int       `json:"id,omitempty"`
	ClientID                         int       `json:"client_id"`
	Name                             string    `json:"name"`
	Code                             string    `json:"code,omitempty"`
	Active                           bool      `json:"active,omitempty"`
	Billable                         bool      `json:"billable,omitempty"`
	BillBy                           string    `json:"bill_by,omitempty"`
	HourlyRate                       float64   `json:"hourly_rate,omitempty"`
	Budget                           float64   `json:"budget,omitempty"`
	BudgetBy                         string    `json:"budget_by,omitempty"`
	NotifyWhenOverBudget             bool      `json:"notify_when_over_budget,omitempty"`
	OverBudgetNotificationPercentage float64   `json:"over_budget_notification_percentage,omitempty"`
	OverBudgetNotifiedAt             string    `json:"over_budget_notified_at,omitempty"`
	ShowBudgetToAll                  bool      `json:"show_budget_to_all,omitempty"`
	CreatedAt                        time.Time `json:"created_at,omitempty"`
	UpdatedAt                        time.Time `json:"updated_at,omitempty"`
	StartsOn                         string    `json:"starts_on,omitempty"`
	EndsOn                           string    `json:"ends_on,omitempty"`
	Estimate                         float64   `json:"estimate,omitempty"`
	EstimateBy                       string    `json:"estimate_by,omitempty"`
	HintEarliestRecordAt             string    `json:"hint_earliest_record_at,omitempty"`
	HintLatestRecordAt               string    `json:"hint_latest_record_at,omitempty"`
	Notes                            string    `json:"notes,omitempty"`
	CostBudget                       float64   `json:"cost_budget,omitempty"`
	CostBudgetIncludeExpenses        bool      `json:"cost_budget_include_expenses,omitempty"`
}