package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	oai "github.com/cristian-sima/Wisply/models/oai"
	repository "github.com/cristian-sima/Wisply/models/repository"
	ws "github.com/cristian-sima/Wisply/models/ws"
)

var hub *ws.Hub

// CurrentProcesses holds the current Statistics for a repository
var CurrentProcesses = make(map[int]*Process)

// Action represents the state (finish) and the number
type Action struct {
	Finished bool `json:"Finished"`
	Number   int  `json:"Number"`
}

// Process contians information about a process
type Process struct {
	CurrentAction string             `json:"CurrentAction"`
	Actions       map[string]*Action `json:"Actions"`
}

func init() {
	hub = ws.CreateHub()
	go hub.Run()
}

// HarvestController It manages the operations for repository (list, delete, add)
type HarvestController struct {
	AdminController
	Model repository.Model
}

// InitWebsocketConnection Initiats the websocket connection
func (controller *HarvestController) InitWebsocketConnection() {
	controller.TplNames = "site/harvest/init.tpl"
	connection := hub.CreateConnection(controller.Ctx.ResponseWriter, controller.Ctx.Request, controller)
	hub.Register <- connection
	go connection.WritePump()
	connection.ReadPump()
}

// DecideAction decides a certain action for the incoming message
func (controller *HarvestController) DecideAction(message *ws.Message, connection *ws.Connection) {

	model := repository.Model{}
	repository, err := model.NewRepository(strconv.Itoa(message.Repository))

	if err != nil {
		fmt.Println("Not a good id of rep from client!")
		fmt.Println(err)
	} else {
		switch message.Name {
		case "changeRepositoryURL":
			newURL := message.Value.(string)
			controller.ChangeRepositoryBaseURL(repository, newURL)
		case "testURL":
			{
				controller.TestURL(repository)
			}
		case "identify":
			{
				controller.IdenfityRepository(repository)
			}
		case "initialize":
			{
				controller.InitializeRepository(repository)
			}
		case "getCurrentProcess":
			{
				controller.GetCurrentProcess(repository)
			}
		}
	}
}

// ChangeRepositoryBaseURL verifies if an address can be reached
func (controller *HarvestController) ChangeRepositoryBaseURL(repository *repository.Repository, newURL string) {

	if newURL != repository.URL {
		repository.ModifyURL(newURL)
	}

	msg := ws.Message{
		Name:       "RepositoryBaseURLChanged",
		Repository: repository.ID,
		Value:      newURL,
	}

	hub.BroadcastMessage(&msg)
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
		CurrentAction: "",
		Actions:       actions,
	}

	CurrentProcesses[ID] = st

	// get records
	controller.getRecords(repository, func(response *oai.Response) {
		fmt.Println("am terminat")

		msg := ws.Message{
			Name:       "FinishStage",
			Repository: repository.ID,
		}

		hub.BroadcastMessage(&msg)

		// delete init
		delete(CurrentProcesses, ID)

		controller.modifyRepositoryStatus(repository, "ok")
	})
}

// GetCurrentProcess gets all the records
func (controller *HarvestController) GetCurrentProcess(repository *repository.Repository) {

	processObject, _ := CurrentProcesses[repository.ID]

	hub.BroadcastMessage(&ws.Message{
		Name:       "processOnServer",
		Value:      &processObject,
		Repository: repository.ID,
	})
}

// GetRecords gets all the records
func (controller *HarvestController) getRecords(repository *repository.Repository, finishCallback func(*oai.Response)) {

	defer func() {
		// recover from any errro and tell them there was a problem
		err := recover()
		if err != nil {
			type Content struct {
				Info string `json:"Info"`
			}
			content := Content{
				Info: err.(string),
			}
			msg := ws.Message{
				Name:       "Record-Problems",
				Value:      content,
				Repository: repository.ID,
			}

			hub.BroadcastMessage(&msg)

			controller.modifyRepositoryStatus(repository, "problems")
		}
	}()

	request := (&oai.Request{
		BaseURL:        repository.URL,
		From:           "2012-02-09T18:12:54Z",
		Until:          "2012-05-09T18:12:54Z",
		MetadataPrefix: "oai_dc",
	})

	request.HarvestRecords(func(record *oai.Record) {

		ID := repository.ID
		actions := CurrentProcesses[ID].Actions
		recordsAction := actions["records"]
		recordsAction.Number++

		fmt.Println("--> I received a record." + strconv.Itoa(actions["records"].Number))
		/*
			type Content struct {
				Data *Action `json:"Data"`
			}
			content := Content{
				Data: recordsAction,
			}
			msg := ws.Message{
				Name:       "Statistics",
				Value:      content,
				Repository: repository.ID,
			}

			hub.BroadcastMessage(&msg)
		*/
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
	msg := ws.Message{
		Name:       "FinishTestingURL",
		Value:      content,
		Repository: repository.ID,
	}

	hub.BroadcastMessage(&msg)

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
				State bool `json:"State"`
			}
			content := Content{
				State: false,
			}
			msg := ws.Message{
				Name:       "FinishIdentify",
				Value:      content,
				Repository: repository.ID,
			}
			hub.BroadcastMessage(&msg)

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
				State bool          `json:"State"`
				Data  *oai.Response `json:"Data"`
			}
			content := Content{
				State: true,
				Data:  record,
			}
			msg := ws.Message{
				Name:       "FinishIdentify",
				Value:      content,
				Repository: repository.ID,
			}
			controller.modifyRepositoryStatus(repository, "verified")
			hub.BroadcastMessage(&msg)
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
		msg := ws.Message{
			Name:       "RepositoryChangedStatus",
			Value:      content,
			Repository: repository.ID,
		}
		hub.BroadcastMessage(&msg)
	} else {
		panic(err)
	}
}

// ShowPanel shows the panel to collect data from repository
func (controller *HarvestController) ShowPanel() {

	ID := controller.Ctx.Input.Param(":id")

	repository, err := controller.Model.NewRepository(ID)

	if err != nil {
		controller.Abort("databaseError")
	}

	controller.Data["repository"] = repository
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/harvest/init.tpl"
	controller.Layout = "site/admin.tpl"
}
