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
			repository.log("I encoured a problem with validation:")
			fmt.Println(err)
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
			repository.log("I encoured a problem with identification request:")
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
		err := recover()
		if err != nil {
			repository.log("I encoured a problem with the harvest formats request:")
			fmt.Println(err)
			msg := Message{
				Name:  "harvesting-failed",
				Value: "failed",
			}
			repository.notifyManager(&msg)
		}
	}()

	formatRequest := (&oai.Request{
		BaseURL: repository.URL,
		Verb:    "ListMetadataFormats",
	})

	formatRequest.Harvest(func(record *oai.Response) {

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
		err := recover()
		if err != nil {
			repository.log("I encoured a problem with the harvest collections request:")
			fmt.Println(err)
			msg := Message{
				Name:  "harvesting-failed",
				Value: "failed",
			}
			repository.notifyManager(&msg)
		}
	}()

	formatRequest := (&oai.Request{
		BaseURL: repository.URL,
		Verb:    "ListSets",
	})

	formatRequest.Harvest(func(response *oai.Response) {

		var collections []Collection

		remoteCollections := response.ListSets.Set

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

// HarvestRecords receives records
func (repository *EPrintsRepository) HarvestRecords() {
	defer func() {
		err := recover()
		if err != nil {
			repository.log("I encoured a problem with the harvest records request:")
			fmt.Println(err)
			msg := Message{
				Name:  "harvesting-failed",
				Value: "failed",
			}
			repository.notifyManager(&msg)
		}
	}()

	formatRequest := (&oai.Request{
		BaseURL:        repository.URL,
		Verb:           "ListRecords",
		MetadataPrefix: "oai_dc",
	})

	formatRequest.Harvest(func(response *oai.Response) {

		var records []Record

		remoteRecords := response.ListRecords.Records

		for _, record := range remoteRecords {

			record := &OAIRecord{
				Identifier: record.Header.Identifier,
				Datestamp:  record.Header.DateStamp,
				Keys:       repository.getKeys(record.Metadata.Body),
			}
			records = append(records, record)
		}

		result := OAIRecordsResult{
			isOk: true,
			data: records,
		}

		repository.Manager.SaveRecords(&result)
	}, func(resp *oai.Response) {
		repository.Manager.EndRecords()
	})
}

func (repository *EPrintsRepository) notifyManager(message *Message) {
	//repository.Manager.Notify(message)
}

func (repository *EPrintsRepository) log(message interface{}) {
	fmt.Println("<-->  EPrints: " + message.(string))
}
