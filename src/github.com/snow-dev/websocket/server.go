package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

//Configure th upgrader

var upgrader = websocket.Upgrader{}

// Define

type Message struct {
	Email    string `json: "email"`
	Username string `json: "username"`
	Message  string `json:"message" `
}

func main() {
	fmt.Printf("\n\tInitializing websocket\n")

	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening  for incoming chat message
	go message()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
