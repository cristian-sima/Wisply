package admin

import (
	"fmt"
	"strconv"

	harvest "github.com/cristian-sima/Wisply/models/harvest"
	repository "github.com/cristian-sima/Wisply/models/repository"
	ws "github.com/cristian-sima/Wisply/models/ws"
)

var hub *ws.Hub

// CurrentSessions contains information about all the running sessions
var CurrentSessions = make(map[int]*Session)

// Session is an object that contains a process and the user connected to it
type Session struct {
	Connections []*ws.Connection `json:"-"`
	Process     *harvest.Process `json:"Manager"`
}

func (process *Session) addConnection(connection *ws.Connection) {
	process.Connections = append(process.Connections, connection)
}

func init() {
	hub = ws.CreateHub()
	go hub.Run()
}

// HarvestController manages the operations for repository (list, delete, add)
type HarvestController struct {
	Controller
	conduit chan harvest.ProcessMessager
	Model   repository.Model
}

// Prepare starts the thread with channel
func (controller *HarvestController) Prepare() {
	controller.Controller.Prepare()
	controller.conduit = make(chan harvest.ProcessMessager, 100)
	go controller.run()
}

// RecoverProcess tries to recover a process
func (controller *HarvestController) RecoverProcess() {

	ID := controller.Ctx.Input.Param(":id")

	// check if it is running
	intID, _ := strconv.Atoi(ID)
	harvestProcess := harvest.NewProcessByID(intID)

	repID := harvestProcess.GetRepository().ID

	_, processExists := CurrentSessions[repID]

	// the process must be finished
	if !processExists {
		delete(CurrentSessions, repID)

		harvest.RecoverProcess(harvestProcess, controller)
		process := &Session{
			Process: harvestProcess,
		}
		CurrentSessions[repID] = process

		go harvestProcess.Recover()
	}
	controller.TplNames = "site/admin/harvest/init.tpl"
}

// ForceFinishProcess terminates a process in an error state
func (controller *HarvestController) ForceFinishProcess() {

	ID := controller.Ctx.Input.Param(":id")

	// check if it is running
	intID, _ := strconv.Atoi(ID)
	harvestProcess := harvest.NewProcessByID(intID)

	harvestProcess.ForceFinish()

	controller.TplNames = "site/admin/harvest/init.tpl"
}

// GetConduit returns the channel for sending and receiving messages
func (controller *HarvestController) GetConduit() chan harvest.ProcessMessager {
	return controller.conduit
}

// InitWebsocketConnection initiats the websocket connection
func (controller *HarvestController) InitWebsocketConnection() {
	controller.TplNames = "site/admin/harvest/init.tpl"
	connection := hub.CreateConnection(controller.Ctx.ResponseWriter, controller.Ctx.Request, controller)
	hub.Register <- connection
	go connection.WritePump()
	connection.ReadPump()
}

// ShowPanel shows the panel to collect data from repository
func (controller *HarvestController) ShowPanel() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	}
	controller.Data["repository"] = repository
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/admin/harvest/init.tpl"
}

// DecideAction decides a certain action for the incoming message
func (controller *HarvestController) DecideAction(message *ws.Message, connection *ws.Connection) {
	if message.Repository != 0 {
		controller.decideOneRepository(message, connection)
	} else {
		controller.decideManyRepositories(message, connection)
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) decideOneRepository(message *ws.Message, connection *ws.Connection) {
	switch message.Name {
	case "change-url":
		{
			controller.ChangeRepositoryBaseURL(message)
		}
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

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) decideManyRepositories(message *ws.Message, connection *ws.Connection) {
	switch message.Name {
	case "get-all-status":
		{
			controller.SendAllRepositoriesStatus(connection)
		}
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) ChangeRepositoryBaseURL(message *ws.Message) {
	newURL := message.Value.(string)
	repository, _ := repository.NewRepository(strconv.Itoa(message.Repository))
	if newURL != repository.URL {
		repository.ModifyURL(newURL)
	}
	msg := ws.Message{
		Name:       "url-changed",
		Repository: repository.ID,
		Value:      newURL,
	}

	hub.BroadcastMessage(&msg)
}

// CreateNewProcess starts the initializing proccess
func (controller *HarvestController) CreateNewProcess(message *ws.Message, connection *ws.Connection) {

	existingProcess, processExists := CurrentSessions[message.Repository]

	if !processExists {
		ID := message.Repository
		delete(CurrentSessions, ID)

		harvestProcess := harvest.CreateProcess(strconv.Itoa(ID), controller)
		process := &Session{
			Process: harvestProcess,
		}
		process.addConnection(connection)
		CurrentSessions[ID] = process

		go harvestProcess.Start()
	} else {
		existingProcess.addConnection(connection)
	}
}

func (controller *HarvestController) run() {
	fmt.Println("Running the controller!! ...")

	for {
		select {
		case message := <-controller.conduit:
			fmt.Println("controller: I received this name : " + message.GetName())

			session, ok := CurrentSessions[message.GetRepository()]

			if ok {
				switch message.GetName() {
				case "repository-status-changed", "identification-details", "event-notice":
					{
						msg := ConvertToWebsocketMessage(message)
						hub.BroadcastMessage(msg)
					}
					break
				case "harvesting", "verification-finished":
					msg := ConvertToWebsocketMessage(message)
					hub.SendGroupMessage(msg, session.Connections)
					break
				case "process-finished":
					{
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
func (controller *HarvestController) GetProcess(message *ws.Message, connection *ws.Connection) {
	controller.log("Yes it goes")
	process, processExists := CurrentSessions[message.Repository]
	if !processExists {
		controller.log("I do not have any process for " + strconv.Itoa(message.Repository))
	} else {
		controller.log("I have a process for the repository " + strconv.Itoa(message.Repository))
		controller.log("I add a new connection for the repository " + strconv.Itoa(message.Repository) + " process")
		process.addConnection(connection)
		fmt.Println(process)
	}
	hub.SendMessage(&ws.Message{
		Name:       "existing-process-on-server",
		Value:      &process,
		Repository: message.Repository,
	}, connection)

}

func (controller *HarvestController) log(message string) {
	fmt.Println("<-->  Harvest Controller: " + message)
}

// SendAllRepositoriesStatus gets all repositories' status only and sends them to a connection
func (controller *HarvestController) SendAllRepositoriesStatus(connection *ws.Connection) {
	list := controller.Model.GetAllStatus()
	hub.SendMessage(&ws.Message{
		Name:  "repositories-status-list",
		Value: &list,
	}, connection)
}
