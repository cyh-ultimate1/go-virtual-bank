package models

import (
	"govirtualbank/globalFunctions"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Account struct {
	gorm.Model
	UserID              uint
	User                User                 `gorm:"ForeignKey:user_id"`
	AccountTransactions []AccountTransaction `gorm:"ForeignKey:AccountID"`
}

type AccountAmount struct {
	CreditAmount decimal.Decimal `sql:"type:decimal(19,4);"`
	DebitAmount  decimal.Decimal `sql:"type:decimal(19,4);"`
}

func (account *Account) GetBalance() decimal.Decimal {
	var db = globalFunctions.OpenDbConnection()
	defer db.Close()

	var accountAmount AccountAmount
	//db.Raw("SELECT SUM(debit_amount) AS DebitAmount FROM account_transactions").Row().Scan(&accountAmount)
	db.Raw("SELECT SUM(debit_amount) AS debit_amount, SUM(credit_amount) AS credit_amount FROM account_transactions WHERE account_id = ?", account.ID).Scan(&accountAmount)

	var balance = accountAmount.DebitAmount.Sub(accountAmount.CreditAmount)
	return balance
}
