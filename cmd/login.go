/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/SudipKatwal/go-chat-app-cobra/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var loginUser helpers.LoginUser

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login user",
	Long:  `Login user with email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		systemLogin()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func promptGetInput(pc string) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func systemLogin() {
	email := promptGetInput("Enter Email: ")

	password := promptGetInput("Enter Password: ")

	user, err := helpers.Login(email, password)

	if err != nil {
		fmt.Println("Email or Password is incorrect. Try again")
		systemLogin()
	}
	loginUser = user
	fmt.Println("Login: ", loginUser)
	chat()
}

func chat() {

	client()

	// fmt.Println("Following is the List of online users: ")

}

func client() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	opts.Query["user"] = loginUser.Email
	opts.Query["pwd"] = loginUser.Password
	uri := "http://localhost:9090/socket.io/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		fmt.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		fmt.Printf("on error\n")
	})

	client.On("connection", func() {
		fmt.Println("Client ID: ")
	})
	client.On("message", func(msg string) {
		fmt.Println("Received: ", msg)
	})
	client.On("disconnection", func() {
		fmt.Printf("on disconnect\n")
	})

	fmt.Println("Chat room started: ")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("chat", command)
		fmt.Printf("Send: %v\n", command)
	}
}
