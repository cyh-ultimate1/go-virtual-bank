package controllers

import (
	"govirtualbank/globalFunctions"
	"govirtualbank/models"
)

func GetAccountByUserName(firstname string, lastname string) models.Account {
	var db = globalFunctions.OpenDbConnection()
	defer db.Close()

	var account models.Account
	db.Raw("SELECT accounts.* FROM accounts INNER JOIN [users] AS us ON accounts.user_id = us.id WHERE us.first_name = ? AND us.last_name = ?", firstname, lastname).Scan(&account)

	return account
}
