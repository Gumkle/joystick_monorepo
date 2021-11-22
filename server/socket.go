package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var counter uint8 = 0
var upgrader = websocket.Upgrader{} // use default options

// osobny endpoint dla gry
func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// czekamy na message, dekodujemy unique id, jest timeout
	player, ok :=c.engine.GetFromGlobalList(unique_id)
	if !ok {
		// wywal error
	}
	room := findRoomWith(player)

	// jezeli unique id sie zagdza
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		if messageType != websocket.BinaryMessage {
			// wywal error
		}
		player.SendMessage(message)
		//room.channel <- message
		// wiadomosc ok


		// sprawdzamy czy binarny
		// jzeli bin to dekodujemy
		// jezeli nie to continue
		log.Printf("Received: %s from conn %d", message, id)
	}
}

func prepareSocket() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}
