package websockets

import (
	"encoding/json"
	"fmt"
)

// Message represents a message from client to server
type Message struct {
	Name       string      `json:"Name"`
	Repository int         `json:"Repository"`
	Value      interface{} `json:"Value"`
}

func broadcastMessage(msg *Message) {
	fmt.Println("--> I broadcast this message: ")
	fmt.Println(msg)
	jsonMsg, _ := json.Marshal(&msg)
	currentHub.broadcast <- jsonMsg
}
