package remote

import (
	"fmt"

	oai "github.com/cristian-sima/Wisply/models/oai"
	local "github.com/cristian-sima/Wisply/models/repository"
)

func NewOAIRepository(local *local.Repository) *OAIRepository {
	oai := &OAIRepository{
		Local: local,
	}
	return oai
}

// OAIRepository encapsulates the functionality for a remote repository using OAI format
type OAIRepository struct {
	Repository
	Local *local.Repository
}

// StartProcess starts the process
func (remote *OAIRepository) StartProcess() {
	remote.notifyController(&Message{
		Value: "The process starts",
	})
	remote.validate()
}

func (remote *OAIRepository) GetLocalRepository() *local.Repository {
	return remote.Local
}

// It receives local notifications
func (remote *OAIRepository) notify(notification *Message) {
	fmt.Println("<--> OAIRepository received notification ")
	fmt.Println(notification)
	switch notification.Name {
	case "verification-finished":
		if notification.Value == "failed" {
			remote.notifyController(notification)
		} else {
			fmt.Println("The validation passed")
		}
		break
	}
}

func (remote *OAIRepository) changeStatus(newStatus string) {
	remote.notifyController(&Message{
		Name:  "status-changed",
		Value: newStatus,
	})
}

func (remote *OAIRepository) notifyController(message *Message) {
	message.Repository = remote.Local.ID
	remote.Controller.Notify(message)
}

// Validate checks if the details are good for a repository is good for OAI format
func (remote *OAIRepository) validate() {
	remote.changeStatus("verifying")
	defer func() {
		err := recover()
		if err != nil {
			msg := Message{
				Name:  "verification-finished",
				Value: "failed",
			}
			remote.changeStatus("verification-failed")
			remote.notify(&msg)
		}
	}()

	request := (&oai.Request{
		BaseURL: remote.Local.URL,
		Verb:    "Identify",
	})

	request.Harvest(func(record *oai.Response) {
		msg := Message{
			Name:  "verification-finished",
			Value: "succeeded",
		}
		remote.changeStatus("verified")
		remote.notify(&msg)
	}, func(resp *oai.Response) {
	})
}
