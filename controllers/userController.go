package controllers

import (
	"errors"
	"fmt"
	"govirtualbank/globalFunctions"
	"govirtualbank/models"
	"log"
)

func GetAllUsers() []models.User {
	var db = globalFunctions.OpenDbConnection()

	defer db.Close()

	var users []models.User
	db.Find(&users)
	return users
}

func GetUserByFirstName(firstname string) models.User {
	var db = globalFunctions.OpenDbConnection()

	defer db.Close()

	var user models.User
	db.Where("first_name = ?", firstname).Find(&user)
	return user
}

func GetUserByName(firstname string, lastname string) models.User {
	var db = globalFunctions.OpenDbConnection()

	defer db.Close()

	var user models.User
	// db.Where("first_name = @firstname AND last_name = @lastname", map[string]interface{}{"firstname": firstname, "lastname": lastname}).Find(&user)
	db.Where("first_name = ? AND last_name = ?", firstname, lastname).Find(&user)

	return user
}

func RegisterUser(register models.APIUserRegister) error {
	var db = globalFunctions.OpenDbConnection()
	defer db.Close()

	var user models.User
	db.Where("first_name = ? AND last_name = ?", register.FirstName, register.LastName).Find(&user)
	if user.ID > 0 {
		return errors.New("Existing user with same name found.")
	}

	hashedPassword, err := globalFunctions.HashPassword(register.Password)
	if err != nil {
		log.Fatal("HashPassword Error: " + err.Error())
	}

	fmt.Println("Creating user...")
	user1 := models.User{FirstName: register.FirstName, LastName: register.LastName, Password: hashedPassword}
	db.Create(&user1)
	account1 := models.Account{}
	account1.UserID = user1.ID
	db.Create(&account1)
	user1.AccountID = account1.ID
	db.Save(&user1)
	return nil
}
