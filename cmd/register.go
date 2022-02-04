/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/SudipKatwal/go-chat-app-cobra/helpers"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register Here: ",
	Long:  `Register new user`,
	Run: func(cmd *cobra.Command, args []string) {
		registerUser()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

func registerUser() {
	name := promptGetInput("Enter Name: ")

	email := promptGetInput("Enter Email: ")

	password := promptGetInput("Enter Password: ")

	helpers.Register(name, email, password)

	systemLogin()
}
