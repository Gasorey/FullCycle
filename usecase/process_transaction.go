package usecase

import (
	"time"

	"github.com/gasorey/fullcycle/domain"
	"github.com/gasorey/fullcycle/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{
		TransactionRepository: transactionRepository,
	}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDTO dto.Transaction) (transaction domain.Transaction, err error) {
	creditCard := u.completeCreditCard(transactionDTO)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.completeTransaction(transactionDTO, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.SaveTransaction(transaction, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (u UseCaseTransaction) completeCreditCard(transactionDTO dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDTO.Name
	creditCard.Number = transactionDTO.Number
	creditCard.ExpirationMonth = transactionDTO.ExpirationMonth
	creditCard.ExpirationYear = transactionDTO.ExpirationYear
	creditCard.CVV = transactionDTO.CVV
	return creditCard
}

func (u UseCaseTransaction) completeTransaction(transactionDTO dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.CreditCardID = cc.ID
	transaction.Amount = transactionDTO.Amount
	transaction.Description = transactionDTO.Description
	transaction.Store = transactionDTO.Store
	transaction.CreatedAt = time.Now()
	return transaction
}
