package models

import "time"

type Transaction struct {
	ID              int `gorm:"primaryKey"`
	AccountID       string
	BankID          string
	Amount          int
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
