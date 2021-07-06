package usecase

import (
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

}
