package admin

import (
	"fmt"
	"strconv"

	harvest "github.com/cristian-sima/Wisply/models/harvest"
	repository "github.com/cristian-sima/Wisply/models/repository"
	ws "github.com/cristian-sima/Wisply/models/ws"
)

var hub *ws.Hub

// CurrentProcesses holds the current Statistics for a repository
var CurrentProcesses = make(map[int]*Process)

// Process contians information about a process
type Process struct {
	Connections []*ws.Connection `json:"-"`
	Manager     *harvest.Manager `json:"Manager"`
}

func (process *Process) addConnection(connection *ws.Connection) {
	process.Connections = append(process.Connections, connection)
}

func init() {
	hub = ws.CreateHub()
	go hub.Run()
}

// HarvestController manages the operations for repository (list, delete, add)
type HarvestController struct {
	Controller
	Model repository.Model
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
			controller.StartProcess(message, connection)
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

// StartProcess starts the initializing proccess
func (controller *HarvestController) StartProcess(message *ws.Message, connection *ws.Connection) {

	_, processExists := CurrentProcesses[message.Repository]

	if !processExists {
		ID := message.Repository
		delete(CurrentProcesses, ID)

		harvestManager := harvest.NewManager(strconv.Itoa(ID), controller)
		process := &Process{
			Manager: harvestManager,
		}
		process.addConnection(connection)
		CurrentProcesses[ID] = process
		harvestManager.StartProcess()
	}
}

// Notify is called by a harvest repository with a message
func (controller *HarvestController) Notify(message *harvest.Message) {
	process, ok := CurrentProcesses[message.Repository]

	fmt.Println("<-->  Harvest Controller: The controller has received this message:")
	fmt.Println(message)
	if ok {
		switch message.Name {
		case "status-changed", "identification-details":
			{
				msg := ConvertToWebsocketMessage(message)
				hub.BroadcastMessage(msg)
			}
			break
		case "verification-finished":
			if message.Value == "failed" {
				msg := ConvertToWebsocketMessage(message)
				hub.SendGroupMessage(msg, process.Connections)
				delete(CurrentProcesses, message.Repository)
			}
		case "harvesting":
			msg := ConvertToWebsocketMessage(message)
			hub.SendGroupMessage(msg, process.Connections)
			break
		case "delete-process":
			{
				delete(CurrentProcesses, message.Repository)
			}
			break
		}
	}
}

// ConvertToWebsocketMessage converts a harvest message to a websocket one
func ConvertToWebsocketMessage(old *harvest.Message) *ws.Message {
	newMessage := &ws.Message{
		Name:       old.Name,
		Value:      old.Value,
		Repository: old.Repository,
	}
	return newMessage
}

// GetProcess sends the current process on the server for a repository
func (controller *HarvestController) GetProcess(message *ws.Message, connection *ws.Connection) {
	processObject, _ := CurrentProcesses[message.Repository]
	hub.SendMessage(&ws.Message{
		Name:       "existing-process-on-server",
		Value:      &processObject,
		Repository: message.Repository,
	}, connection)
}

// SendAllRepositoriesStatus gets all repositories' status only and sends them to a connection
func (controller *HarvestController) SendAllRepositoriesStatus(connection *ws.Connection) {
	list := controller.Model.GetAllStatus()
	hub.SendMessage(&ws.Message{
		Name:  "repositories-status-list",
		Value: &list,
	}, connection)
}
