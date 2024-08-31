package models

import "time"

type InMemAccountDB map[int]Account

type Account struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`

	Transaction []Transaction `json:"transactions"`
}

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type Transaction struct {
	ID              int
	TransactionType TransactionType
	Amount          float64
	TransTimeStamp  time.Time
}

func GetTime() time.Time {
	return time.Now()
}

type Error struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
