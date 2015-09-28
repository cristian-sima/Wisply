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
		fmt.Println("Probleeeeeeeeeeeeeeeeeeeeeeeeeeem")
		fmt.Println(err)
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
		remoteIdentity := record.Identify

		identify := OAIIdentification{
			Name:              remoteIdentity.RepositoryName,
			URL:               remoteIdentity.BaseURL,
			Protocol:          remoteIdentity.ProtocolVersion,
			AdminEmails:       remoteIdentity.AdminEmail,
			EarliestDatestamp: remoteIdentity.EarliestDatestamp,
			RecordPolicy:      remoteIdentity.DeletedRecord,
			Granularity:       remoteIdentity.Granularity,
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
