package models

import "time"

type Payment struct {
	ID            int64     `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Email         string    `json:"email,omitempty"`
	Method        string    `json:"method,omitempty"`
	TransactionId string    `json:"transaction_id,omitempty"`
	Amount        int64     `json:"amount,omitempty"`
	PlanId        int       `json:"plan_id,omitempty"`
	PayAt         time.Time `json:"pay_at,omitempty"`
}
