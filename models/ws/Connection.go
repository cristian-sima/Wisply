package websockets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WebController defines a controller which can receive messages from the hub
type WebController interface {
	DecideAction(message *Message)
}

// Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	hub *Hub

	controller WebController
}

// readPump pumps messages from the websocket connection to the hub.
func (connection *Connection) ReadPump() {
	defer func() {
		connection.hub.Unregister <- connection
		connection.ws.Close()
	}()
	connection.ws.SetReadLimit(MaxMessageSize)
	connection.ws.SetReadDeadline(time.Now().Add(ReadWaitPeriod))
	for {
		operation, r, err := connection.ws.NextReader()
		if err != nil {
			break
		}
		switch operation {
		case websocket.PongMessage:
			connection.ws.SetReadDeadline(time.Now().Add(ReadWaitPeriod))
		case websocket.TextMessage:
			messageByte, err := ioutil.ReadAll(r)
			if err != nil {
				break
			}
			var msg Message
			json.Unmarshal(messageByte, &msg)

			fmt.Println("<-- I received the message: ")
			fmt.Println(msg)

			connection.controller.DecideAction(&msg)

		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (connection *Connection) WritePump() {
	ticker := time.NewTicker(PingInterval)
	defer func() {
		ticker.Stop()
		connection.ws.Close()
	}()
	for {
		select {
		case message, ok := <-connection.send:
			if !ok {
				connection.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := connection.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := connection.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// write writes a message with the given opCode and payload.
func (connection *Connection) write(opCode int, payload []byte) error {
	connection.ws.SetWriteDeadline(time.Now().Add(WriteWaitPeriod))
	return connection.ws.WriteMessage(opCode, payload)
}
