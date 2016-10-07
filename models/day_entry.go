package models

import "time"

type DayEntry struct {
	ID               int       `json:"id"`
	Notes            string    `json:"notes"`
	SpentAt          string    `json:"spent_at"`
	Hours            float64   `json:"hours"`
	UserID           int       `json:"user_id"`
	ProjectID        int       `json:"project_id"`
	TaskID           int       `json:"task_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AdjustmentRecord bool      `json:"adjustment_record"`
	TimerStartedAt   string    `json:"timer_started_at"`
	IsClosed         bool      `json:"is_closed"`
	IsBilled         bool      `json:"is_billed"`
}
