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
		fmt.Println("Defer identification")
		err := recover()
		if err != nil {
			fmt.Println("Identification problem")
			fmt.Println(err)
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

// HarvestFormats receives the formats
func (repository *EPrintsRepository) HarvestFormats() {
	defer func() {
		fmt.Println("Defer formats")
		err := recover()
		if err != nil {
			fmt.Println("Formats problem")
			msg := Message{
				Name:  "harvesting-failed",
				Value: "failed",
			}
			fmt.Println("failed")
			repository.notifyManager(&msg)
		}
	}()

	fmt.Println("Trimit pentru formats")
	formatRequest := (&oai.Request{
		BaseURL: repository.URL,
		Verb:    "ListMetadataFormats",
	})

	formatRequest.Harvest(func(record *oai.Response) {
		fmt.Println("le-am luat")

		var formats []Formater

		remoteFormats := record.ListMetadataFormats.MetadataFormat

		for _, format := range remoteFormats {
			format := &OAIFormat{
				Prefix:    format.MetadataPrefix,
				Namespace: format.MetadataNamespace,
				Schema:    format.Schema,
			}
			formats = append(formats, format)
		}

		result := OAIFormatResult{
			isOk: true,
			data: formats,
		}

		repository.Manager.SaveFormats(&result)
	}, func(resp *oai.Response) {
		repository.Manager.EndFormats()
	})
}

// HarvestCollections receives the sets alias collections
func (repository *EPrintsRepository) HarvestCollections() {
	defer func() {
		fmt.Println("Defer collections")
		err := recover()
		if err != nil {
			fmt.Println("!!!! collections problem")
			fmt.Println(err)
			msg := Message{
				Name:  "harvesting-failed",
				Value: "failed",
			}
			repository.notifyManager(&msg)
		}
	}()

	fmt.Println("Trimit pentru collections")
	formatRequest := (&oai.Request{
		BaseURL: repository.URL,
		Verb:    "ListSets",
	})

	formatRequest.Harvest(func(response *oai.Response) {
		fmt.Println("le-am luat collections")

		var collections []Collection

		fmt.Println("in fata")
		remoteCollections := response.ListSets.Set

		fmt.Println(remoteCollections)
		fmt.Println("dupa")
		for _, collection := range remoteCollections {
			collection := &OAICollection{
				Name: collection.SetName,
				Spec: collection.SetSpec,
			}
			collections = append(collections, collection)
		}

		result := OAICollectionResult{
			isOk: true,
			data: collections,
		}

		repository.Manager.SaveCollections(&result)
	}, func(resp *oai.Response) {
		repository.Manager.EndCollections()
	})
}

func (repository *EPrintsRepository) notifyManager(message *Message) {
	repository.Manager.Notify(message)
}
