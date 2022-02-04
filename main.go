/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"net/http"

	"github.com/SudipKatwal/go-chat-app-cobra/cmd"
	"github.com/SudipKatwal/go-chat-app-cobra/helpers"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	helpers.Connect()

	cmd.Execute()
	server := socketio.NewServer(nil)

	server.OnConnect("", func(s socketio.Conn) error {
		fmt.Println("Connected: ", s.ID())
		s.Rooms()
		s.Join("chat")
		return nil
	})

	server.OnEvent("", "chat", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom("", "chat", "message", msg)
	})

	server.OnError("", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	fmt.Println("Serving at localhost:9090...")
	http.ListenAndServe(":9090", nil)
}
