package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Hub manages the connections
type Hub struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	Register chan *Connection

	// Unregister requests from connections.
	Unregister chan *Connection
}

// CreateNewConnection creates a new ws connection
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

func (hub *Hub) SendMessage(message *Message, connection *Connection) {
	fmt.Println("--> I send the next message to one connection: ")
	fmt.Println(message)
	jsonMsg, _ := json.Marshal(&message)
	select {
	case connection.send <- jsonMsg:
	default:
		close(connection.send)
		delete(hub.connections, connection)
	}
}

func (hub *Hub) BroadcastMessage(message *Message) {
	fmt.Println("--> I broadcast this message: ")
	fmt.Println(message)
	jsonMsg, _ := json.Marshal(&message)
	hub.broadcast <- jsonMsg
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
				select {
				case connection.send <- message:
				default:
					close(connection.send)
					delete(hub.connections, connection)
				}
			}
		}
	}
}
