package websockets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/gorilla/websocket"
)

// Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (connection *Connection) readPump() {
	defer func() {
		currentHub.unregister <- connection
		connection.ws.Close()
	}()
	connection.ws.SetReadLimit(maxMessageSize)
	connection.ws.SetReadDeadline(time.Now().Add(readWait))
	for {
		op, r, err := connection.ws.NextReader()
		if err != nil {
			break
		}
		switch op {
		case websocket.PongMessage:
			connection.ws.SetReadDeadline(time.Now().Add(readWait))
		case websocket.TextMessage:
			messageByte, err := ioutil.ReadAll(r)
			if err != nil {
				break
			}

			fmt.Println("<-- I received the message: ")

			var msg Message

			json.Unmarshal(messageByte, &msg)
			connection.chooseAction(msg)
			fmt.Println(msg)
		}
	}
}

func (connection *Connection) chooseAction(msg Message) {

	model := repository.Model{}
	rep, err := model.NewRepository(strconv.Itoa(msg.Repository))

	if err != nil {
		fmt.Println("Not a good id of rep from client!")
		fmt.Println(err)
	} else {

		switch msg.Name {
		case "changeRepositoryURL":
			newURL := msg.Value.(string)
			connection.controller.ChangeRepositoryBaseURL(rep, newURL)
		case "testURL":
			{
				connection.controller.TestURL(rep)
			}
		case "identify":
			{
				connection.controller.IdenfityRepository(rep)
			}
		case "initialize":
			{
				connection.controller.InitializeRepository(rep)
			}
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (connection *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
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
	connection.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return connection.ws.WriteMessage(opCode, payload)
}
