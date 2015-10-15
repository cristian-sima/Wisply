package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Hub represents the object which managers the message. It sends messages to the connections and it notifies the controller if there are new messages
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
		hub.log("I have a problem with handshaking")
		fmt.Println(err)
	} else if err != nil {
		hub.log("I have a different problem")
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

// SendMessage sends only one message to a single connection
func (hub *Hub) SendMessage(message *Message, connection *Connection) {
	// hub.log("--> Hub: I send to a single connection the message: ")
	hub.sendWebsocket(message, connection)
}

// SendGroupMessage sends a message to a GROUP of connections
func (hub *Hub) SendGroupMessage(message *Message, group []*Connection) {
	// hub.log("I broadcast to a GROUP of " + strconv.Itoa(len(group)) + " connections, this message: ")
	for _, connection := range group {
		hub.sendWebsocket(message, connection)
	}
}

// BroadcastMessage sends a message to ALL the connection from the hub
func (hub *Hub) BroadcastMessage(message *Message) {
	// 	hub.log("I broadcast to ALL " + strconv.Itoa(len(hub.connections)) + " connections, this message: ")
	if hub.broadcast != nil {
		hub.broadcast <- message
	}
}

// Run starts the main chanel. It registers, unregisters and broadcasts messages
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

// log prints a message in a nice format
func (hub *Hub) log(message string) {
	// fmt.Println("--> Hub: " + message)
}

// sendWebsocket converts the message to a websocket and sends it to the connection
func (hub *Hub) sendWebsocket(message *Message, connection *Connection) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Hub error ignored:")
		}
	}()

	ws, err := json.Marshal(&message)
	if err != nil {
		hub.log("I got an error when I tried to compress the websocket message into json at sendWebsocket[ at line 94]. Here is the error:")
		fmt.Println(err)
	}
	select {
	case connection.send <- ws:
		//	hub.log("Websocket sent")
	default:
		fmt.Println("a ajuns aici")
		close(connection.send)
		delete(hub.connections, connection)
	}
}
