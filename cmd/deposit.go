/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"govirtualbank/controllers"

	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deposit called")

		// var db = globalFunctions.OpenDbConnection()
		// defer db.Close()

		depositAmount, _ := cmd.Flags().GetFloat32("depositAmount")
		destinationFirstName, _ := cmd.Flags().GetString("destinationFirstName")
		destinationLastName, _ := cmd.Flags().GetString("destinationLastName")
		fmt.Println(destinationFirstName)
		fmt.Println(destinationLastName)

		controllers.DepositToAccount(decimal.NewFromFloat32(depositAmount), destinationFirstName, destinationLastName)

	},
}

func init() {
	rootCmd.AddCommand(depositCmd)

	depositCmd.Flags().Float32("depositAmount", 0.00, "enter deposit amount")
	depositCmd.Flags().String("destinationFirstName", "", "enter destination firstname")
	depositCmd.Flags().String("destinationLastName", "", "enter destination lastname")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depositCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// depositCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
