package services

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lonmarsDev/banking-backend/models"
)

type BankService struct {
	// save in memory
	AccountsDb models.InMemAccountDB
	mutex      sync.Mutex
}

func NewBankService() BankService {
	return BankService{
		AccountsDb: models.InMemAccountDB{},
	}
}

func (b *BankService) CreateAccount(name string, balance float64) models.Account {
	account := models.Account{
		ID:      len(b.AccountsDb) + 1,
		Name:    name,
		Balance: balance,
	}
	b.AccountsDb[account.ID] = account
	return account
}

func (b *BankService) GetAccount(id int) (models.Account, error) {
	account, ok := b.AccountsDb[id]
	if !ok {
		return models.Account{}, fmt.Errorf("account with id %d not found", id)
	}
	return account, nil

}

func (b *BankService) Deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	account, err := b.GetAccount(id)
	if err != nil {
		return errors.New("account not found")
	}
	// insert transaction
	transaction := models.Transaction{
		ID:              len(account.Transaction) + 1,
		TransactionType: models.Deposit,
		Amount:          amount,
		TransTimeStamp:  models.GetTime(),
	}
	account.Transaction = append(account.Transaction, transaction)
	// lock the account to prevent race condition
	b.mutex.Lock()
	defer b.mutex.Unlock()
	account.Balance += amount
	b.AccountsDb[id] = account
	return nil
}

func (b *BankService) Withdraw(id int, amount float64) error {
	account, err := b.GetAccount(id)
	if err != nil {
		return errors.New("account not found")
	}
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}
	// insert transaction
	transaction := models.Transaction{
		ID:              len(account.Transaction) + 1,
		TransactionType: models.Withdraw,
		Amount:          amount,
		TransTimeStamp:  models.GetTime(),
	}
	account.Transaction = append(account.Transaction, transaction)
	// lock the account to prevent race condition
	b.mutex.Lock()
	defer b.mutex.Unlock()
	account.Balance -= amount
	b.AccountsDb[id] = account
	return nil
}

func (b *BankService) GetTransactions(id int) ([]models.Transaction, error) {
	account, err := b.GetAccount(id)
	if err != nil {
		return []models.Transaction{}, errors.New("account not found")
	}
	return account.Transaction, nil
}
