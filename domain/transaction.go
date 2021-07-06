package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardID string
	CreatedAt    time.Time
}

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) (err error)
	GetCreditCard(creditCard CreditCard) (CreditCard, err error)
	CreateCreditCard(creditCard CreditCard) (err error)
}

func NewTransaction() *Transaction {
	transaction := &Transaction{}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	return transaction
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {

	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}