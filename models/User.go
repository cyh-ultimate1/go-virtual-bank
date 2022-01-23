package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Password  string `json:"-"`
	AccountID uint
	//Account   Account `gorm:"ForeignKey:account_id"`
}

type APIUser struct {
	FirstName string
	LastName  string
}

type APIUserRegister struct {
	FirstName string
	LastName  string
	Password  string
}
