/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"govirtualbank/constants"
	"govirtualbank/models"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// migrateToDbCmd represents the migrateToDb command
var migrateToDbCmd = &cobra.Command{
	Use:   "migrateToDb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateToDb called")

		input, _ := cmd.Flags().GetString("dbAction")

		fmt.Println(input)

		connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			constants.Server, constants.User, constants.Password, constants.Port, constants.DatabaseName)
		db, err := gorm.Open("mssql", connectionString)

		if err != nil {
			log.Fatal("Failed to create connection pool. Error: " + err.Error())
		}

		fmt.Println("connected")
		defer db.Close()

		migrate(input, db)
	},
}

func init() {
	rootCmd.AddCommand(migrateToDbCmd)

	migrateToDbCmd.Flags().String("dbAction", "all", "Help message for toggle")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateToDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateToDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func migrate(input string, db *gorm.DB) {
	fmt.Println("Migrating models...")
	if input == "User" {
		db.AutoMigrate(&models.User{})
	} else if input == "Account" {
		db.AutoMigrate(&models.Account{})
	} else if input == "AccountTransaction" {

		db.AutoMigrate(&models.AccountTransaction{})
	} else if input == "TransactionType" {
		db.AutoMigrate(&models.TransactionType{})
	} else {
		db.AutoMigrate(&models.User{}, &models.Account{}, &models.AccountTransaction{}, &models.TransactionType{})
		db.Model(&models.Account{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
		//db.Model(&models.User{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")
	}
}
