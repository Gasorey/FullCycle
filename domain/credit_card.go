package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth int64
	ExpirationYear  int64
	CVV             int64
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func NewCreditCard() *CreditCard {
	creditCard := &CreditCard{}
	creditCard.ID = uuid.NewV4().String()
	creditCard.CreatedAt = time.Now()

	return creditCard
}
