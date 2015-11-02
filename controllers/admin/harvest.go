package admin

import (
	"fmt"
	"strconv"

	harvest "github.com/cristian-sima/Wisply/models/harvest"
	repository "github.com/cristian-sima/Wisply/models/repository"
	ws "github.com/cristian-sima/Wisply/models/ws"
)

var (
	hub     *ws.Hub
	conduit = make(chan harvest.ProcessMessager, 100)
)

// CurrentSessions contains information about all the running sessions
var CurrentSessions = make(map[int]*session)

// Session is an object that contains a process and the user connected to it
type session struct {
	Connections []*ws.Connection `json:"-"`
	Process     *harvest.Process `json:"Manager"`
}

func (process *session) addConnection(connection *ws.Connection) {
	process.Connections = append(process.Connections, connection)
}

func init() {
	hub = ws.CreateHub()
	go hub.Run()
	go run()
}

// Harvest manages the operations for repository
type Harvest struct {
	Controller
}

// RecoverProcess tries to recover a process
func (controller *Harvest) RecoverProcess() {
	ID := controller.Ctx.Input.Param(":id")
	// check if it is running
	intID, _ := strconv.Atoi(ID)
	harvestProcess := harvest.NewProcessByID(intID)
	repID := harvestProcess.GetRepository().ID
	delete(CurrentSessions, repID)
	harvest.RecoverProcess(harvestProcess, controller)
	process := &session{
		Process: harvestProcess,
	}
	CurrentSessions[repID] = process
	go harvestProcess.Recover()

	controller.TplNames = "site/admin/harvest/init.tpl"
}

// ForceFinishProcess terminates a process in an error state
func (controller *Harvest) ForceFinishProcess() {
	ID := controller.Ctx.Input.Param(":id")
	// check if it is running
	intID, _ := strconv.Atoi(ID)
	harvestProcess := harvest.NewProcessByID(intID)
	harvestProcess.ForceFinish()
	controller.TplNames = "site/admin/harvest/init.tpl"
}

// GetConduit returns the channel for sending and receiving messages
func (controller *Harvest) GetConduit() chan harvest.ProcessMessager {
	if conduit == nil {
		panic("conduit nil stop")
	}
	return conduit
}

// InitWebsocketConnection initiats the websocket connection
func (controller *Harvest) InitWebsocketConnection() {
	controller.TplNames = "site/admin/harvest/init.tpl"
	response := controller.Ctx.ResponseWriter
	request := controller.Ctx.Request
	connection := hub.CreateConnection(response, request, controller)
	hub.Register <- connection
	go connection.WritePump()
	connection.ReadPump()
}

// ShowPanel shows the panel to collect data from repository
func (controller *Harvest) ShowPanel() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["repository"] = repository
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/admin/harvest/init.tpl"
}

// DecideAction decides a certain action for the incoming message
func (controller *Harvest) DecideAction(message *ws.Message, connection *ws.Connection) {
	if message.Repository != 0 {
		controller.decideOneRepository(message, connection)
	} else {
		controller.decideManyRepositories(message, connection)
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *Harvest) decideOneRepository(message *ws.Message, connection *ws.Connection) {
	switch message.Name {
	case "start-progress":
		{
			controller.CreateNewProcess(message, connection)
		}
	case "get-current-progress":
		{
			controller.GetProcess(message, connection)
		}
	}
}

// CreateNewProcess starts the initializing proccess
func (controller *Harvest) CreateNewProcess(message *ws.Message, connection *ws.Connection) {
	_, processExists := CurrentSessions[message.Repository]
	if !processExists {
		ID := message.Repository
		delete(CurrentSessions, ID)
		IDString := strconv.Itoa(ID)
		harvestProcess := harvest.CreateProcess(IDString, controller)
		process := &session{
			Process: harvestProcess,
		}
		process.addConnection(connection)
		CurrentSessions[ID] = process
		go harvestProcess.Start()
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *Harvest) decideManyRepositories(message *ws.Message, connection *ws.Connection) {
	switch message.Name {
	case "get-all-status":
		{
			controller.SendAllRepositoriesStatus(connection)
		}
	}
}

func run() {
	for {
		select {
		case message := <-conduit:
			session, ok := CurrentSessions[message.GetRepository()]
			if ok {
				switch message.GetName() {
				case "repository-status-changed", "identification-details", "event-notice":
					{
						msg := ConvertToWebsocketMessage(message)
						hub.BroadcastMessage(msg)
					}
					break
				case "harvest-update":
					if session == nil {
						panic("session nil")
					}
					if session.Connections == nil {
						panic("session connections nil")
					}
					msg := ConvertToWebsocketMessage(message)
					hub.SendGroupMessage(msg, session.Connections)
					break
				case "process-finished":
					{
						msg := ConvertToWebsocketMessage(message)
						hub.SendGroupMessage(msg, session.Connections)
						delete(CurrentSessions, message.GetRepository())
					}
					break
				}
			}
		}
	}
}

// ConvertToWebsocketMessage converts a harvest message to a websocket one
func ConvertToWebsocketMessage(old harvest.ProcessMessager) *ws.Message {
	newMessage := &ws.Message{
		Name:       old.GetName(),
		Value:      old.GetValue(),
		Repository: old.GetRepository(),
	}
	return newMessage
}

// GetProcess sends the current process on the server for a repository
func (controller *Harvest) GetProcess(message *ws.Message, connection *ws.Connection) {
	process, processExists := CurrentSessions[message.Repository]
	if !processExists {
		IDString := strconv.Itoa(message.Repository)
		errorMessage := "I do not have any process for " + IDString
		controller.log(errorMessage)
	}
	hub.SendMessage(&ws.Message{
		Name:       "existing-process-on-server",
		Value:      &process,
		Repository: message.Repository,
	}, connection)

}

func (controller *Harvest) log(message string) {
	fmt.Println("<-->  Harvest Controller: " + message)
}

// SendAllRepositoriesStatus gets all repositories' status
// It only sends messages to a connection
func (controller *Harvest) SendAllRepositoriesStatus(connection *ws.Connection) {
	list := repository.GetAllStatus()
	hub.SendMessage(&ws.Message{
		Name:  "repositories-status-list",
		Value: &list,
	}, connection)
}
