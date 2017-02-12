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

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we  close the connection when the function returns
	defer ws.Close()

	// register our client
	clients[ws] = true

	for {
		var msg Message

		// Read in new message as JSON and map it to a message object
		err := ws.ReadJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessage() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for clients := range clients {
			err := clients.WriteJSON(msg)
			if err != nill {
				log.Printf("error: %v", err)
				clients.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	fmt.Printf("\n\tInitializing websocket\n")

	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening  for incoming chat message
	go handleMessage()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
