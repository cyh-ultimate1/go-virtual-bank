package globalFunctions

import (
	"fmt"
	"govirtualbank/constants"
	"log"

	"github.com/jinzhu/gorm"
)

func OpenDbConnection() *gorm.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		constants.Server, constants.User, constants.Password, constants.Port, constants.DatabaseName)
	db, err := gorm.Open("mssql", connectionString)
	if err != nil {
		log.Fatal("Failed to create connection pool. Error: " + err.Error())
	}
	fmt.Println("connected")
	gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")
	//defer db.Close()

	return db
}
