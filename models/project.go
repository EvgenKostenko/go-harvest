package models

import "time"

type Project struct {
	ID                               int       `json:"id"`
	ClientID                         int       `json:"client_id"`
	Name                             string    `json:"name"`
	Code                             string    `json:"code"`
	Active                           bool      `json:"active"`
	Billable                         bool      `json:"billable"`
	BillBy                           string    `json:"bill_by"`
	HourlyRate                       float64   `json:"hourly_rate"`
	Budget                           float64   `json:"budget"`
	BudgetBy                         string    `json:"budget_by"`
	NotifyWhenOverBudget             bool      `json:"notify_when_over_budget"`
	OverBudgetNotificationPercentage float64   `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             string `json:"over_budget_notified_at"`
	ShowBudgetToAll                  bool      `json:"show_budget_to_all"`
	CreatedAt                        time.Time `json:"created_at"`
	UpdatedAt                        time.Time `json:"updated_at"`
	StartsOn                         string    `json:"starts_on"`
	EndsOn                           string    `json:"ends_on"`
	Estimate                         float64   `json:"estimate"`
	EstimateBy                       string    `json:"estimate_by"`
	HintEarliestRecordAt             string    `json:"hint_earliest_record_at"`
	HintLatestRecordAt               string    `json:"hint_latest_record_at"`
	Notes                            string    `json:"notes"`
	CostBudget                       float64   `json:"cost_budget"`
	CostBudgetIncludeExpenses        bool      `json:"cost_budget_include_expenses"`
}
