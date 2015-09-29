package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

// Hub manages the connections
type Hub struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan *Message

	// Register requests from the connections.
	Register chan *Connection

	// Unregister requests from connections.
	Unregister chan *Connection
}

// CreateConnection creates a new ws connection
func (hub *Hub) CreateConnection(response http.ResponseWriter, request *http.Request, controller WebController) *Connection {

	ws, err := upgrader.Upgrade(response, request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		fmt.Println("I have a problem")
	} else if err != nil {
		fmt.Println("I have a problem")
		fmt.Println(err)
	}
	connection := &Connection{
		send:       make(chan []byte, 256),
		ws:         ws,
		hub:        hub,
		controller: controller,
	}
	return connection
}

// SendMessage sends a message to ONE connection
func (hub *Hub) SendMessage(message *Message, connection *Connection) {
	hub.log("--> Hub: I send to a single connection the message: ")
	fmt.Println(message)
	hub.sendWebsocket(message, connection)
}

// BroadcastMessage sends a message to ALL the connection from the hub
func (hub *Hub) BroadcastMessage(message *Message) {
	hub.log("I broadcast to ALL " + strconv.Itoa(len(hub.connections)) + " connections, this message: ")
	fmt.Println(message)
	hub.broadcast <- message
}

// SendGroupMessage sends a message to a GROUP of connections
func (hub *Hub) SendGroupMessage(message *Message, group []*Connection) {
	hub.log("I broadcast to a GROUP of " + strconv.Itoa(len(group)) + " connections, this message: ")
	fmt.Println(message)
	for _, connection := range group {
		hub.sendWebsocket(message, connection)
	}
}

// Run starts the hub
func (hub *Hub) Run() {
	for {
		select {
		case connection := <-hub.Register:
			hub.connections[connection] = true
		case connection := <-hub.Unregister:
			delete(hub.connections, connection)
			close(connection.send)
		case message := <-hub.broadcast:
			for connection := range hub.connections {
				hub.sendWebsocket(message, connection)
			}
		}
	}
}

func (hub *Hub) log(message string) {
	fmt.Println("--> Hub: " + message)
}

func (hub *Hub) sendWebsocket(message *Message, connection *Connection) {
	ws, _ := json.Marshal(&message)
	select {
	case connection.send <- ws:
		fmt.Println("go: websocket sent")
	default:
		close(connection.send)
		delete(hub.connections, connection)
	}
}
