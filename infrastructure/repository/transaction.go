package repository

import (
	"database/sql"

	"github.com/gasorey/fullcycle/domain"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{
		db: db,
	}
}

func (t *TransactionRepositoryDB) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) (err error) {

	stmt, err := t.db.Prepare(`
		INSERT INTO
			transactions(
				id,
				credit_card_id,
				amount,
				status,
				description,
				store,
				created_at
			)
			VALUES (
				?,
				?,
				?,
				?,
				?,
				?,
				?
			)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardID,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = t.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) updateBalance(creditCard domain.CreditCard) (err error) {
	_, err = t.db.Exec(`UPDATE credit_cards set balance = ? where id = ?`, creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepositoryDB) CreateCreditCard(creditCard domain.CreditCard) (err error) {
	stmt, err := t.db.Prepare(`
		INSERT INTO
			credit_cards(
				id,
				name,
				number,
				expiration_month,
				expiration_year,
				cvv,
				balance,
				limit,
				created_at
			)	
			VALUES(
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?
			)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
		creditCard.CreatedAt,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
