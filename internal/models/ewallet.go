package models

import "time"

type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" gorm:"column:user_id;unique" validate:"required" `
	Balance   float64   `json:"balance" gorm:"column:balance;type:decimal(15,2)" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                     int
	UserID                 int     `gorm:"column:user_id;type:int" validate:"required"`
	Amount                 float64 `gorm:"column:balance;type:decimal(15,2)" validate:"required"`
	EwalletTransactionType string  `gorm:"column:wallet_transaction_type;type:enum('CREDIT','DEBIT')"`
	Refrence               string  `gorm:"column:refrence;type:varchar(100);type:varchar(100)"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}
