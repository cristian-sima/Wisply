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

// Validate checks if the details are good for a repository is good for EPrints format
func (repository *EPrintsRepository) Validate() {
	defer func() {
		err := recover()
		if err != nil {
			msg := Message{
				Name:  "verification-finished",
				Value: "failed",
			}
			repository.notifyManager(&msg)
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
		repository.notifyManager(&msg)
	}, func(resp *oai.Response) {
	})
}

// HarvestIdentification checks if the details are good for a repository is good for EPrints format
func (repository *EPrintsRepository) HarvestIdentification() {

	defer func() {
		err := recover()
		if err != nil {
			msg := Message{
				Name:  "identification-harvested-finished",
				Value: "failed",
			}
			repository.notifyManager(&msg)
		}
	}()

	request := (&oai.Request{
		BaseURL: repository.URL,
		Verb:    "Identify",
	})

	request.Harvest(func(record *oai.Response) {
		fmt.Println("The harvested record")
		remoteIdentity := record.Identify
		identify := OAIIdentification{
			name:              remoteIdentity.RepositoryName,
			url:               remoteIdentity.BaseURL,
			protocol:          remoteIdentity.ProtocolVersion,
			adminEmails:       remoteIdentity.AdminEmail,
			earliestDatestamp: remoteIdentity.EarliestDatestamp,
			recordPolicy:      remoteIdentity.DeletedRecord,
			granularity:       remoteIdentity.Granularity,
		}

		result := OAIIdentificationResult{
			isOk: true,
			data: &identify,
		}

		repository.Manager.SaveIdentification(&result)
	}, func(resp *oai.Response) {
	})
}

func (repository *EPrintsRepository) notifyManager(message *Message) {
	repository.Manager.Notify(message)
}
