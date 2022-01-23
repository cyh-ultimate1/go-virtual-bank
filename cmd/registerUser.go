/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"govirtualbank/constants"
	"govirtualbank/globalFunctions"
	"govirtualbank/models"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// registerUserCmd represents the registerUser command
var registerUserCmd = &cobra.Command{
	Use:   "registerUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("registerUser called")

		firstname, _ := cmd.Flags().GetString("firstname")
		lastname, _ := cmd.Flags().GetString("lastname")
		password, _ := cmd.Flags().GetString("password")

		connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
			constants.Server, constants.User, constants.Password, constants.Port, constants.DatabaseName)
		db, err := gorm.Open("mssql", connectionString)
		if err != nil {
			log.Fatal("Failed to create connection pool. Error: " + err.Error())
		}
		fmt.Println("connected")
		gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")
		defer db.Close()

		hashedPassword, err := globalFunctions.HashPassword(password)
		if err != nil {
			log.Fatal("HashPassword Error: " + err.Error())
		}

		fmt.Println("Creating user...")
		user1 := models.User{FirstName: firstname, LastName: lastname, Password: hashedPassword}
		db.Create(&user1)
		account1 := models.Account{}
		account1.UserID = user1.ID
		db.Create(&account1)
		user1.AccountID = account1.ID
		db.Save(&user1)
	},
}

func init() {
	rootCmd.AddCommand(registerUserCmd)

	registerUserCmd.Flags().String("firstname", "", "enter firstname")
	registerUserCmd.Flags().String("lastname", "", "enter lastname")
	registerUserCmd.Flags().String("password", "", "enter password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
