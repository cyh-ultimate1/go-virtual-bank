package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type AccountTransaction struct {
	gorm.Model
	CreditAmount      decimal.Decimal `sql:"type:decimal(19,4);"`
	DebitAmount       decimal.Decimal `sql:"type:decimal(19,4);"`
	Remarks           string
	AccountID         uint
	TransactionTypeID uint
	TransactionType   TransactionType
}

type APIAccountTransactionRequest struct {
	Amount              decimal.Decimal `sql:"type:decimal(19,4);"`
	Remarks             string
	ReceipientFirstName string
	ReceipientLastName  string
	SourceFirstName     string
	SourceLastName      string
}
