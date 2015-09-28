package harvest

import (
	"fmt"

	oai "github.com/cristian-sima/Wisply/models/oai"
)

// EPrintsRepository encapsulates the functionality for a repository repository using OAI format
type EPrintsRepository struct {
	RemoteRepository
	URL string
}

// StartProcess starts the process
func (repository *EPrintsRepository) Start() {
	repository.notifyManager(&Message{
		Value: "The process starts",
	})
	repository.validate()
}

// It receives local notifications
func (repository *EPrintsRepository) notify(notification *Message) {
	fmt.Println("<--> EPrintsRepository received notification ")
	fmt.Println(notification)
	switch notification.Name {
	case "verification-finished":
		if notification.Value == "failed" {
			repository.notifyManager(notification)
		} else {
			fmt.Println("The validation passed")
		}
		break
	}
}

func (repository *EPrintsRepository) changeStatus(newStatus string) {
	repository.notifyManager(&Message{
		Name:  "status-changed",
		Value: newStatus,
	})
}

func (repository *EPrintsRepository) notifyManager(message *Message) {
	repository.Manager.Notify(message)
}

// Validate checks if the details are good for a repository is good for OAI format
func (repository *EPrintsRepository) validate() {
	repository.changeStatus("verifying")
	defer func() {
		err := recover()
		if err != nil {
			msg := Message{
				Name:  "verification-finished",
				Value: "failed",
			}
			repository.changeStatus("verification-failed")
			repository.notify(&msg)
		}
	}()

	request := (&oai.Request{
		BaseURL: repository.URL,
		Verb:    "Identify",
	})

	request.Harvest(func(record *oai.Response) {
		msg := Message{
			Name:  "verification-finished",
			Value: "succeeded",
		}
		repository.changeStatus("verified")
		repository.notify(&msg)
	}, func(resp *oai.Response) {
	})
}
