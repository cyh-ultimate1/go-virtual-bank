/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"govirtualbank/controllers"

	"github.com/spf13/cobra"
)

// getBalanceCmd represents the getBalance command
var getBalanceCmd = &cobra.Command{
	Use:   "getBalance",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getBalance called")

		destinationFirstName, _ := cmd.Flags().GetString("destinationFirstName")
		destinationLastName, _ := cmd.Flags().GetString("destinationLastName")
		var user = controllers.GetUserByName(destinationFirstName, destinationLastName)
		fmt.Println(user.FirstName)
		//fmt.Println(user.Account.GetBalance())

		var account = controllers.GetAccountByUserName(destinationFirstName, destinationLastName)
		if account.ID <= 0 {
			return
		}

		fmt.Println(account.GetBalance())

		// var db = globalFunctions.OpenDbConnection()
		// defer db.Close()
		// //var amounts uint
		// //var tx models.AccountTransaction
		// var aa tempStruct
		// db.Model(&models.AccountTransaction{}).Select("sum(id) AS idTotal").Scan(&aa)
		// //db.Raw("SELECT TOP 1 account_id AS DebitAmount FROM account_transactions").Scan(&amounts)
		// fmt.Println(aa.idTotal)

		// var count uint
		// //db.Table("account_transactions").Select("count(id)").Scan(&count)
		// db.Table("account_transactions").Select("sum(transaction_type_id)").Row().Scan(&count)
		// fmt.Println(count)
	},
}

type tempStruct struct {
	idTotal uint
}

func init() {
	rootCmd.AddCommand(getBalanceCmd)

	getBalanceCmd.Flags().String("destinationFirstName", "", "enter destination firstname")
	getBalanceCmd.Flags().String("destinationLastName", "", "enter destination lastname")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
