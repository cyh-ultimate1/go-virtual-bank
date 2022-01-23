package main

import (
	"context"
	"database/sql"
	"fmt"
	"govirtualbank/apiControllers"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var db *sql.DB

func main() {
	////Uncomment below to use CLI interface.
	// cmd.Execute()

	////Below section will use API interface.
	apiControllers.HandleRequests()
	//keep running until next line key press
	fmt.Scanln()
	fmt.Println("done")
}

// Gets and prints SQL Server version
func SelectVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
