package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	oai "github.com/cristian-sima/Wisply/models/oai"
	websocket "github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next message from the client.
	readWait = 6000 * time.Second

	// Send pings to client with this period. Must be less than readWait.
	pingPeriod = (readWait * 9) / 10

	// Maximum message size allowed from client.
	maxMessageSize = 512
)

func init() {
	go h.run()
}

var h = &hub{
	broadcast:   make(chan []byte, maxMessageSize),
	register:    make(chan *connection, 1),
	unregister:  make(chan *connection, 1),
	connections: make(map[*connection]bool),
}

type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		}
	}
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	username string

	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	controller *RepositoryController
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(readWait))
	for {
		op, r, err := c.ws.NextReader()
		if err != nil {
			break
		}
		switch op {
		case websocket.PongMessage:
			c.ws.SetReadDeadline(time.Now().Add(readWait))
		case websocket.TextMessage:
			messageByte, err := ioutil.ReadAll(r)
			if err != nil {
				break
			}
			type Message struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}

			var msg Message

			json.Unmarshal(messageByte, &msg)
			c.chooseAction(msg.Name, msg.Value)
			fmt.Println(msg)
		}
	}
}

func (c *connection) chooseAction(name, value string) {
	switch name {
	case "testURL":
		{
			c.controller.TestURL(value)
		}
	case "identify":
		{
			c.controller.IdenfityRepository()
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// write writes a message with the given opCode and payload.
func (c *connection) write(opCode int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(opCode, payload)
}

// JSONMessage represents a message sent from server to browser
// It contains 2 fields: the name and the content
type JSONMessage struct {
	Name    string      `json:"name"`
	Content interface{} `json:"content"`
}

// RepositoryController It manages the operations for sources (list, delete, add)
type RepositoryController struct {
	AdminController
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// WebsocketConnection Initiats the websocket connection
func (controller *RepositoryController) WebsocketConnection() {
	// Print the OAI Response object to stdout

	ws, err := upgrader.Upgrade(controller.Ctx.ResponseWriter, controller.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(controller.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws, username: "lalalla", controller: controller}
	h.register <- c
	go c.writePump()
	c.readPump()
}

// TestURL verifies if an address can be reached
func (controller *RepositoryController) TestURL(address string) {

	var isOk bool

	request, err := http.Get(address)
	fmt.Println(request)
	if request == nil || err != nil {
		isOk = false
	} else if http.StatusOK != request.StatusCode {
		isOk = false
	}

	// change in database

	if !isOk {
		h.broadcast <- []byte("{\"name\": \"FinishTestingURL\", \"content\" :\"false\"}")
		fmt.Println("nu e bun")
	} else {
		h.broadcast <- []byte("{\"name\": \"FinishTestingURL\", \"content\":\"true\"}")
		fmt.Println("e bun")
	}

}

// IdenfityRepository requests an identification
func (controller *RepositoryController) IdenfityRepository() {

	fmt.Println("This is the test")
	req := (&oai.Request{
		BaseURL: "http://eprints.aston.ac.uk/cgi/oai2",
		Verb:    "Identify",
	})

	req.Harvest(func(record *oai.Response) {

		type Content struct {
			State bool          `json:"state"`
			Data  *oai.Response `json:"data"`
		}

		content := Content{
			State: true,
			Data:  record,
		}

		msg := JSONMessage{
			Name:    "FinishIdentify",
			Content: content,
		}
		jsonMsg, _ := json.Marshal(&msg)
		fmt.Println("maica")
		fmt.Println(jsonMsg)
		s := string(jsonMsg[:])
		fmt.Println(s)
		h.broadcast <- jsonMsg
	})
}

// ShowPanel shows the panel to collect data from repository
func (controller *RepositoryController) ShowPanel() {
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/repository/index.tpl"
	controller.Layout = "site/admin.tpl"
}
