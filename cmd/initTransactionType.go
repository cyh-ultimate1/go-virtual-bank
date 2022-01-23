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

// initTransactionTypeCmd represents the initTransactionType command
var initTransactionTypeCmd = &cobra.Command{
	Use:   "initTransactionType",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initTransactionType called")

		//var db = globalFunctions.OpenDbConnection()

		connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			constants.Server, constants.User, constants.Password, constants.Port, constants.DatabaseName)
		db, err := gorm.Open("mssql", connectionString)
		if err != nil {
			log.Fatal("Failed to create connection pool. Error: " + err.Error())
		}
		fmt.Println("connected")
		gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")
		defer db.Close()
		initTransactionType(db)
	},
}

func init() {
	rootCmd.AddCommand(initTransactionTypeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initTransactionTypeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initTransactionTypeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initTransactionType(db *gorm.DB) {
	fmt.Println("Initializing transaction types....")
	db.Create(&models.TransactionType{TypeName: constants.DepositText})
	db.Create(&models.TransactionType{TypeName: constants.WithdrawText})
	db.Create(&models.TransactionType{TypeName: constants.TransferText})
	fmt.Println("Initializing transaction types done.")
}
