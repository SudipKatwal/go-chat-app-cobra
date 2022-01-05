/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/SudipKatwal/go-chat-app-cobra/cmd"
)

func main() {
	database.connect()

	cmd.Execute()
}
