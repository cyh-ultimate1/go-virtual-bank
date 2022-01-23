package controllers

import (
	"errors"
	"fmt"
	"govirtualbank/constants"
	"govirtualbank/globalFunctions"
	"govirtualbank/models"

	"github.com/shopspring/decimal"
)

func DepositToAccount(depositAmount decimal.Decimal, receipientFirstName string, receipientLastName string) error {
	var db = globalFunctions.OpenDbConnection()

	defer db.Close()

	fmt.Println("depositing....")
	destUser := models.User{}
	account := models.Account{}
	transactionType := models.TransactionType{}

	db.Where("first_name = ? AND last_name = ?", receipientFirstName, receipientLastName).Find(&destUser)
	if destUser.ID <= 0 {
		return errors.New("Recipient User not found")
	}

	db.Where("user_id = ?", destUser.ID).Find(&account)
	if account.ID <= 0 {
		return errors.New("Recipient User account not found")
	}

	db.Where("type_name = ?", constants.DepositText).Find(&transactionType)
	if (models.TransactionType{}) == transactionType {
		return errors.New("Recipient User transaction type not found")
	}

	db.Create(&models.AccountTransaction{DebitAmount: depositAmount, AccountID: account.ID, TransactionTypeID: transactionType.ID})

	fmt.Println("deposit done.")
	return nil

}

func WithdrawFromAccount(amount decimal.Decimal, sourceFirstName string, sourceLastName string) error {
	var db = globalFunctions.OpenDbConnection()

	defer db.Close()

	fmt.Println("withdrawing....")
	destUser := models.User{}
	account := models.Account{}
	transactionType := models.TransactionType{}

	db.Where("first_name = ? AND last_name = ?", sourceFirstName, sourceLastName).Find(&destUser)
	if destUser.ID <= 0 {
		return errors.New("Recipient User not found")
	}

	db.Where("user_id = ?", destUser.ID).Find(&account)
	if account.ID <= 0 {
		return errors.New("Recipient User account not found")
	}

	if account.GetBalance().Sub(amount).LessThan(decimal.NewFromInt(0)) {
		return errors.New("Account balance insufficient.")
	}

	db.Where("type_name = ?", constants.WithdrawText).Find(&transactionType)
	if (models.TransactionType{}) == transactionType {
		return errors.New("Recipient User transaction type not found")
	}

	db.Create(&models.AccountTransaction{CreditAmount: amount, AccountID: account.ID, TransactionTypeID: transactionType.ID})

	fmt.Println("Withdraw done.")
	return nil

}

func TransferToAccount(amount decimal.Decimal, sourceFirstName string, sourceLastName string, receipientFirstName string, receipientLastName string) error {
	var db = globalFunctions.OpenDbConnection()
	defer db.Close()

	fmt.Println("transferring....")
	destUser := models.User{}
	sourceUser := models.User{}
	account := models.Account{}
	sourceUserAccount := models.Account{}
	transactionType := models.TransactionType{}

	db.Where("first_name = ? AND last_name = ?", receipientFirstName, receipientLastName).Find(&destUser)
	if destUser.ID <= 0 {
		return errors.New("Recipient User not found")
	}

	db.Where("first_name = ? AND last_name = ?", sourceFirstName, sourceLastName).Find(&sourceUser)
	if sourceUser.ID <= 0 {
		return errors.New("Source User not found")
	}

	db.Where("user_id = ?", destUser.ID).Find(&account)
	if account.ID <= 0 {
		return errors.New("Recipient User account not found")
	}

	db.Where("user_id = ?", sourceUser.ID).Find(&sourceUserAccount)
	if sourceUserAccount.ID <= 0 {
		return errors.New("Source User account not found")
	}

	if sourceUserAccount.GetBalance().Sub(amount).LessThan(decimal.NewFromInt(0)) {
		return errors.New("Account balance insufficient.")
	}

	db.Where("type_name = ?", constants.WithdrawText).Find(&transactionType)
	if (models.TransactionType{}) == transactionType {
		return errors.New("Recipient User transaction type not found")
	}

	db.Create(&models.AccountTransaction{CreditAmount: amount, AccountID: sourceUserAccount.ID, TransactionTypeID: transactionType.ID})
	db.Create(&models.AccountTransaction{DebitAmount: amount, AccountID: account.ID, TransactionTypeID: transactionType.ID})

	fmt.Println("Transfer done.")
	return nil
}
