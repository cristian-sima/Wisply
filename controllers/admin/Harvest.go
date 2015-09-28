package admin

import (
	"fmt"
	"strconv"

	remote "github.com/cristian-sima/Wisply/models/remote"
	repository "github.com/cristian-sima/Wisply/models/repository"
	ws "github.com/cristian-sima/Wisply/models/ws"
)

var hub *ws.Hub

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

const (
	// TESTING is the id of testing stage
	TESTING = 3
	// IDENTIFYING is the id of identifying stage
	IDENTIFYING = 4
)

// CurrentProcesses holds the current Statistics for a repository
var CurrentProcesses = make(map[int]*Process)

// Action represents the state (finish) and the number
type Action struct {
	Finished bool `json:"Finished"`
	Number   int  `json:"Number"`
}

// Process contians information about a process
type Process struct {
	CurrentAction int                `json:"CurrentAction"`
	Actions       map[string]*Action `json:"Actions"`
	Connections   []*ws.Connection   `json:"-"`
	Remote        remote.Standard    `json:"-"`
}

func (process *Process) addConnection(connection *ws.Connection) {
	process.Connections = append(process.Connections, connection)
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
	repository, err := repository.NewRepository(strconv.Itoa(message.Repository))
	if err != nil {
		fmt.Println("No repository with that id.")
	} else {
		switch message.Name {
		case "change-url":
			newURL := message.Value.(string)
			controller.ChangeRepositoryBaseURL(repository, newURL)
		case "start-progress":
			{
				controller.StartProcess(repository, connection)
			}
		case "get-current-progress":
			{
				controller.GetCurrentProcess(repository, connection)
			}
		}
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) decideManyRepositories(message *ws.Message, connection *ws.Connection) {
	switch message.Name {
	case "get-all-status":
		{
			controller.GetAllRepositoriesStatus(connection)
		}
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) ChangeRepositoryBaseURL(repository *repository.Repository, newURL string) {
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
func (controller *HarvestController) StartProcess(local *repository.Repository, connection *ws.Connection) {

	ID := local.ID

	// delete any previous
	delete(CurrentProcesses, ID)

	actions := map[string]*Action{
		"records": {
			Number:   0,
			Finished: false,
		},
	}

	remoteRepository := remote.NewOAIRepository(local)
	remoteRepository.SetController(controller)

	// create a new empty one
	process := &Process{
		CurrentAction: TESTING,
		Actions:       actions,
		Remote:        remoteRepository,
	}

	process.addConnection(connection)

	CurrentProcesses[ID] = process

	remoteRepository.StartProcess()

	/*if repository.Status != "verified" {
		controller.startVerifyRepository(repository, connection)
	} else {
		controller.startInit(repository)
	}*/
}

// Notify is called by a remote repository with a message
func (controller *HarvestController) Notify(message *remote.Message) {
	process, ok := CurrentProcesses[message.Repository]

	fmt.Println("--> Harvest Controller: The controller has received this message:")
	fmt.Println(message)
	if ok {
		switch message.Name {
		case "status-changed":
			{
				process.Remote.GetLocalRepository().ModifyStatus(message.Value)
				msg := ConvertToWebsocketMessage(message)
				hub.BroadcastMessage(msg)
			}
		case "verification-finished":
			if message.Value == "failed" {
				msg := ConvertToWebsocketMessage(message)
				hub.BroadcastMessage(msg)
				delete(CurrentProcesses, message.Repository)
			}
			break
		}
	}
}

// ConvertToWebsocketMessage converts a remote message to a websocket one
func ConvertToWebsocketMessage(old *remote.Message) *ws.Message {
	newMessage := &ws.Message{
		Name:       old.Name,
		Value:      old.Value,
		Repository: old.Repository,
	}
	return newMessage
}

// GetCurrentProcess gets all the records
func (controller *HarvestController) GetCurrentProcess(repository *repository.Repository, connection *ws.Connection) {
	processObject, _ := CurrentProcesses[repository.ID]
	hub.SendMessage(&ws.Message{
		Name:       "existing-process-on-server",
		Value:      &processObject,
		Repository: repository.ID,
	}, connection)
}

// GetAllRepositoriesStatus gets all repositories' status only
func (controller *HarvestController) GetAllRepositoriesStatus(connection *ws.Connection) {
	list := controller.Model.GetAllStatus()
	hub.SendMessage(&ws.Message{
		Name:  "repositories-status-list",
		Value: &list,
	}, connection)
}

// NotifyProcessChanged tells the connection the current process
func (controller *HarvestController) NotifyProcessChanged(repository *repository.Repository, connection *ws.Connection) {
	processObject, _ := CurrentProcesses[repository.ID]
	hub.SendMessage(&ws.Message{
		Name:       "ProcessUpdated",
		Value:      &processObject,
		Repository: repository.ID,
	}, connection)
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
