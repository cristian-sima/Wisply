package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	oai "github.com/cristian-sima/Wisply/models/oai"
	repository "github.com/cristian-sima/Wisply/models/repository"
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

type Action struct {
	Finished bool `json:"Finished"`
	Number   int  `json:"Number"`
}

type Process struct {
	Current string             `json:"CurrentAction"`
	Actions map[string]*Action `json:"Actions"`
}

// CurrentProcesses holds the current Statistics for a repository
var CurrentProcesses = make(map[int]*Process)

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

	controller *HarvestController
}

// Message represents a message from client to server
type Message struct {
	Name       string      `json:"Name"`
	Repository int         `json:"Repository"`
	Value      interface{} `json:"Value"`
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

			fmt.Println("<-- I received the message: ")

			var msg Message

			json.Unmarshal(messageByte, &msg)
			c.chooseAction(msg)
			fmt.Println(msg)
		}
	}
}

func (c *connection) chooseAction(msg Message) {

	model := repository.Model{}
	rep, err := model.NewRepository(strconv.Itoa(msg.Repository))

	if err != nil {
		fmt.Println("Not a good id of rep from client!")
		fmt.Println(err)
	} else {

		switch msg.Name {
		case "changeRepositoryURL":
			newURL := msg.Value.(string)
			c.controller.ChangeRepositoryBaseURL(rep, newURL)
		case "testURL":
			{
				c.controller.TestURL(rep)
			}
		case "identify":
			{
				c.controller.IdenfityRepository(rep)
			}
		case "initialize":
			{
				c.controller.InitializeRepository(rep)
			}
		}
	}
}

func broadcastMessage(msg *Message) {
	fmt.Println("--> I broadcast this message: ")
	fmt.Println(msg)
	jsonMsg, _ := json.Marshal(&msg)
	h.broadcast <- jsonMsg
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

// HarvestController It manages the operations for repository (list, delete, add)
type HarvestController struct {
	AdminController
	Model repository.Model
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// InitWebsocketConnection Initiats the websocket connection
func (controller *HarvestController) InitWebsocketConnection() {
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

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) ChangeRepositoryBaseURL(repository *repository.Repository, newURL string) {

	if newURL != repository.URL {
		repository.ModifyURL(newURL)
	}

	msg := Message{
		Name:       "RepositoryBaseURLChanged",
		Repository: repository.ID,
		Value:      newURL,
	}

	broadcastMessage(&msg)
}

// InitializeRepository starts the initializing proccess
func (controller *HarvestController) InitializeRepository(repository *repository.Repository) {
	if repository.Status != "verified" {
		fmt.Println("Can't start init, because the rep is not in 'verified' state. It is in:" + repository.Status)
	} else {
		controller.modifyRepositoryStatus(repository, "initializing")
		controller.startInit(repository)
	}
}

func (controller *HarvestController) startInit(repository *repository.Repository) {

	ID := repository.ID

	// delete any previous
	delete(CurrentProcesses, ID)

	actions := map[string]*Action{
		"records": &Action{
			Number:   0,
			Finished: false,
		},
	}

	// create a new empty one
	st := &Process{
		Current: "",
		Actions: actions,
	}

	CurrentProcesses[ID] = st

	// get records
	controller.getRecords(repository, func(response *oai.Response) {
		fmt.Println("am terminat")

		msg := Message{
			Name:       "FinishStage",
			Repository: repository.ID,
		}
		broadcastMessage(&msg)

		// delete init
		delete(CurrentProcesses, ID)

		controller.modifyRepositoryStatus(repository, "ok")
	})
}

// GetRecords gets all the records
func (controller *HarvestController) getRecords(repository *repository.Repository, finishCallback func(*oai.Response)) {

	defer func() {
		// recover from any errro and tell them there was a problem
		err := recover()
		if err != nil {
			type Content struct {
				Info string `json:"info"`
			}
			content := Content{
				Info: err.(string),
			}
			msg := Message{
				Name:       "Record-Problems",
				Value:      content,
				Repository: repository.ID,
			}
			broadcastMessage(&msg)

			controller.modifyRepositoryStatus(repository, "problems")
		}
	}()

	request := (&oai.Request{
		BaseURL:        repository.URL,
		From:           "2012-02-09T18:12:54Z",
		Until:          "2014-04-09T18:12:54Z",
		MetadataPrefix: "oai_dc",
	})

	request.HarvestRecords(func(record *oai.Record) {

		ID := repository.ID
		actions := CurrentProcesses[ID].Actions
		recordsAction := actions["records"]
		recordsAction.Number++

		fmt.Println("--> I received a record." + strconv.Itoa(actions["records"].Number))

		type Content struct {
			Data *Action `json:"Data"`
		}
		content := Content{
			Data: recordsAction,
		}
		msg := Message{
			Name:       "Statistics",
			Value:      content,
			Repository: repository.ID,
		}

		broadcastMessage(&msg)
	}, finishCallback)

}

// TestURL verifies if an address can be reached
func (controller *HarvestController) TestURL(repository *repository.Repository) {

	controller.modifyRepositoryStatus(repository, "verifying")

	var isOk bool

	isOk = true

	request, err := http.Get(repository.URL)
	if request == nil || err != nil {
		isOk = false
	} else if http.StatusOK != request.StatusCode {
		isOk = false
	}

	type Content struct {
		State bool `json:"IsValid"`
	}
	content := Content{
		State: isOk,
	}
	msg := Message{
		Name:       "FinishTestingURL",
		Value:      content,
		Repository: repository.ID,
	}

	broadcastMessage(&msg)

	if !isOk {
		controller.modifyRepositoryStatus(repository, "verification-failed")
	}
}

// IdenfityRepository requests an identification
func (controller *HarvestController) IdenfityRepository(repository *repository.Repository) {

	defer func() {
		// recover from any errro and tell them there was a problem
		err := recover()
		if err != nil {
			type Content struct {
				State bool `json:"state"`
			}
			content := Content{
				State: false,
			}
			msg := Message{
				Name:       "FinishIdentify",
				Value:      content,
				Repository: repository.ID,
			}
			broadcastMessage(&msg)

			controller.modifyRepositoryStatus(repository, "verification-failed")
		}
	}()

	if repository.Status != "verifying" {
		fmt.Println("The repository should be in 'verifying' state")
	} else {
		request := (&oai.Request{
			BaseURL: repository.URL,
			Verb:    "Identify",
		})

		request.Harvest(func(record *oai.Response) {

			type Content struct {
				State bool          `json:"state"`
				Data  *oai.Response `json:"data"`
			}
			content := Content{
				State: true,
				Data:  record,
			}
			msg := Message{
				Name:       "FinishIdentify",
				Value:      content,
				Repository: repository.ID,
			}
			controller.modifyRepositoryStatus(repository, "verified")
			broadcastMessage(&msg)
		}, func(resp *oai.Response) {
		})
	}
}

func (controller *HarvestController) modifyRepositoryStatus(repository *repository.Repository, newStatus string) {
	err := repository.ModifyStatus(newStatus)

	if err == nil {
		type Content struct {
			NewStatus string `json:"NewStatus"`
		}
		content := Content{
			NewStatus: repository.Status,
		}
		msg := Message{
			Name:       "RepositoryChangedStatus",
			Value:      content,
			Repository: repository.ID,
		}
		broadcastMessage(&msg)
	} else {
		panic(err)
	}
}

// ShowPanel shows the panel to collect data from repository
func (controller *HarvestController) ShowPanel() {

	var processDetails string

	ID := controller.Ctx.Input.Param(":id")

	repository, err := controller.Model.NewRepository(ID)

	if err != nil {
		controller.Abort("databaseError")
	}

	temp1, _ := strconv.Atoi(ID)
	processObject, hasProcess := CurrentProcesses[temp1]

	if hasProcess {
		temp2, _ := json.Marshal(processObject)
		processDetails = string(temp2)
	}

	controller.Data["hasProcess"] = hasProcess
	controller.Data["currentProcesses"] = processDetails

	controller.Data["repository"] = repository
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/harvest/init.tpl"
	controller.Layout = "site/admin.tpl"
}
