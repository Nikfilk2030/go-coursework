package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type Client struct {
//	username string
//	conn     *websocket.Conn
//}
//
//var clients = make(map[*Client]bool)
//var broadcast = make(chan Message)
//
////var upgrader = websocket.Upgrader{}
//
//type Message struct {
//	Author string `json:"author"`
//	Text   string `json:"text"`
//}

//func handleConnections(w http.ResponseWriter, r *http.Request) {
//	// Upgrade initial GET request to a websocket
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ws.Close()
//
//	// Register our new client
//	client := &Client{conn: ws}
//	clients[client] = true
//
//	for {
//		var msg Message
//		// Read in a new message as JSON and map it to a Message object
//		err := ws.ReadJSON(&msg)
//		if err != nil {
//			log.Printf("error: %v", err)
//			delete(clients, client)
//			break
//		}
//		msg.Username = client.username
//		msg.Time = time.Now()
//		// Send the newly received message to the broadcast channel
//		broadcast <- msg
//	}
//}

//func handleMessages() {
//	for {
//		// Grab the next message from the broadcast channel
//		msg := <-broadcast
//		// Send it out to every client that is currently connected
//		for client := range clients {
//			if client.username == msg.Username {
//				err := client.conn.WriteJSON(msg)
//				if err != nil {
//					log.Printf("error: %v", err)
//					client.conn.Close()
//					delete(clients, client)
//				}
//			}
//		}
//	}
//}

// ////////////////////////////////////////////////////////////////////////

type Message struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

var messages []Message

func main() {
	port := "8080"

	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var message Message
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&message)
			fmt.Println("got message:", message.Author, message.Text)
			if err != nil {
				panic(err)
			}

			messages = append(messages, message)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(message)
			return
		}

		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(messages)
			fmt.Println("sent messages")
			return
		}

		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	})

	// program listens on port
	http.ListenAndServe(":"+port, nil)
}

//////////////////////////////////////////////////////////////////////////

//func main() {
//	// Create a simple file server
//	fs := http.FileServer(http.Dir("./public"))
//	http.Handle("/", fs)
//
//	// Configure websocket route
//	http.HandleFunc("/ws", handleConnections)
//
//	// Start listening for incoming chat posts
//	go handleMessages()
//
//	// Start the server on localhost port 8080 and log any errors
//	log.Println("http server started on :8080")
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}
