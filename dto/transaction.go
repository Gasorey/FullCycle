package dto

import "time"

type Transaction struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int64
	ExpirationYear  int64
	CVV             int64
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}
