package oai

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/harvest/remote/oai/protocol"
	"github.com/cristian-sima/Wisply/models/harvest/remote/repository"
	wisply "github.com/cristian-sima/Wisply/models/wisply/data"
)

// Repository holds a local reposiory, the request and response and a filter
type Repository struct {
	*repository.Repository
	finishToken  string
	request      *protocol.Request
	lastResponse *protocol.Response
	filter       filter
}

// ------------------------- GET

// Test calls the request validate method and returns its result
func (repository *Repository) Test() ([]byte, error) {
	return repository.request.Identify()
}

// Identify returns the identification details
func (repository *Repository) Identify() ([]byte, error) {
	return repository.request.Identify()
}

// ListFormats returns the content of the request which requests for formats
func (repository *Repository) ListFormats() ([]byte, error) {
	return repository.request.GetFormats()
}

// ListCollections returns the content of the request which requests the sets
func (repository *Repository) ListCollections(token string) ([]byte, error) {
	repository.request.Clear()
	repository.request.ResumptionToken = token
	return repository.request.GetSets()
}

// ListRecords returns the content of the request which requests for records
func (repository *Repository) ListRecords(token string) ([]byte, error) {
	repository.request.Clear()
	repository.request.ResumptionToken = token
	return repository.request.GetRecords("oai_dc")
}

// ListIdentifiers returns the content of the request which requests for identifiers
func (repository *Repository) ListIdentifiers(token string) ([]byte, error) {
	repository.request.Clear()
	repository.request.ResumptionToken = token
	return repository.request.GetIdentifiers("oai_dc")
}

// ----------------------------------- Parse

// GetIdentification returns the identification details in Wisply format
func (repository *Repository) GetIdentification(content []byte) (*wisply.Identificationer, error) {

	var idenfitication wisply.Identificationer

	response, err := repository.request.Parse(content)
	if err != nil {
		return &idenfitication, err
	}

	remoteIdentity := response.Identify

	// check a field which must be
	if remoteIdentity.RepositoryName == "" {
		message := "There was a problem getting the fields"
		return &idenfitication, errors.New(message)
	}

	idenfitication = &Identification{
		Name:              remoteIdentity.RepositoryName,
		URL:               remoteIdentity.BaseURL,
		Protocol:          remoteIdentity.ProtocolVersion,
		AdminEmails:       remoteIdentity.AdminEmail,
		EarliestDatestamp: remoteIdentity.EarliestDatestamp,
		RecordPolicy:      remoteIdentity.DeletedRecord,
		Granularity:       remoteIdentity.Granularity,
	}

	return &idenfitication, nil
}

// GetFormats returns the `formats` in Wisply format
func (repository *Repository) GetFormats(content []byte) ([]wisply.Formater, error) {

	var formats []wisply.Formater

	response, err := repository.request.Parse(content)
	if err != nil {
		return formats, err
	}

	remoteFormats := response.ListMetadataFormats.MetadataFormat

	for _, format := range remoteFormats {
		format := Format{
			Prefix:    format.MetadataPrefix,
			Namespace: format.MetadataNamespace,
			Schema:    format.Schema,
		}
		formats = append(formats, &format)
	}

	return formats, nil
}

// GetCollections returns the `collections` in Wisply format
func (repository *Repository) GetCollections(content []byte) ([]wisply.Collectioner, error) {
	var collections []wisply.Collectioner

	response, err := repository.request.Parse(content)
	if err != nil {
		return collections, err
	}

	var getNameFromPath = func(path string) string {
		name := ""
		elements := strings.Split(path, ":")
		name = elements[len(elements)-1]
		name = strings.Replace(name, "=", "-", -1)
		return name
	}

	// cache the last response
	repository.lastResponse = response

	remoteCollections := response.ListSets.Set

	for _, collection := range remoteCollections {

		collectionName := getNameFromPath(collection.SetName)

		collection := &Collection{
			Path: collection.SetName,
			Spec: collection.SetSpec,
			Name: collectionName,
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

// GetRecords returns the `records` in Wisply format
func (repository *Repository) GetRecords(content []byte) ([]wisply.Recorder, error) {
	var records []wisply.Recorder
	response, err := repository.request.Parse(content)
	// cache the last response
	repository.lastResponse = response
	if err != nil {
		return records, err
	}
	remoteRecords := response.ListRecords.Records

	getKeys := func(plainText []byte) (Keys, error) {
		keys := Keys{}
		err := xml.Unmarshal(plainText, &keys)
		if err != nil {
			return keys, err
		}
		return keys, nil
	}

	for _, record := range remoteRecords {
		keys, err := getKeys(record.Metadata.Body)
		if err == nil {
			record := Record{
				Identifier: record.Header.Identifier,
				Datestamp:  record.Header.DateStamp,
				Keys:       keys,
			}
			if repository.filter.isRecordAllowed(record) {
				records = append(records, record)
			}
		}
	}
	if len(remoteRecords) != 0 {
		repository.prepareFinishToken()
	}
	return records, nil
}

// GetIdentifiers returns the `records` in Wisply format
func (repository *Repository) GetIdentifiers(content []byte) ([]wisply.Identifier, error) {
	var identifiers []wisply.Identifier
	response, err := repository.request.Parse(content)
	// cache the last response
	repository.lastResponse = response
	if err != nil {
		return identifiers, err
	}
	remoteIdentifiers := response.ListIdentifiers.Headers

	for _, remoteIdentifier := range remoteIdentifiers {
		identifier := Identifier{
			Identifier: remoteIdentifier.Identifier,
			Datestamp:  remoteIdentifier.DateStamp,
			Spec:       remoteIdentifier.SetSpec,
		}
		if repository.filter.isIdentifierAllowed(identifier) {
			identifiers = append(identifiers, &identifier)
		}
	}
	return identifiers, nil
}

// GetNextPage checks if the last response has a resumption token
func (repository *Repository) GetNextPage() (string, bool) {
	var (
		token    string
		hasToken bool
	)
	if repository.lastResponse == nil {
		return token, false
	}
	hasToken = repository.lastResponse.HasResumptionToken()
	token = repository.lastResponse.GetResumptionToken()
	if !hasToken {
		return token, false
	}
	return token, hasToken
}

// GetFinishToken returns the finishing token
func (repository *Repository) GetFinishToken() string {
	return repository.finishToken
}

func (repository *Repository) prepareFinishToken() {
	resp := repository.lastResponse
	record := resp.ListRecords.Records[len(resp.ListRecords.Records)-1]
	identifier := record.Header.Identifier
	elements := strings.Split(identifier, ":")
	id, _ := strconv.Atoi(elements[2])
	repository.finishToken = "metadataPrefix%3Doai_dc%26offset%3D" + strconv.Itoa(id+1)
}

// IsValidResponse checks if the content is an OAI format
func (repository *Repository) IsValidResponse(content []byte) error {
	return repository.request.IsValidResponse(content)
}

// newRepository returns a new oai repository
func newRepository(rep *repository.Repository) *Repository {
	req := &protocol.Request{
		BaseURL: rep.GetLocalRepository().URL,
	}
	return &Repository{
		Repository: rep,
		request:    req,
		filter:     newFilter(rep.GetLocalRepository().GetFilter()),
	}
}
