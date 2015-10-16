package oai

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/harvest/remote/oai/protocol"
	"github.com/cristian-sima/Wisply/models/harvest/remote/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// EPrintsRepository is an Eprints remote repository
type EPrintsRepository struct {
	*repository.Repository
	finishToken  string
	request      *protocol.Request
	lastResponse *protocol.Response
	filter       filter
}

// ------------------------- GET

// Test calls the request validate method and returns its result
func (repository *EPrintsRepository) Test() ([]byte, error) {
	return repository.request.Identify()
}

// Identify returns the identification details
func (repository *EPrintsRepository) Identify() ([]byte, error) {
	return repository.request.Identify()
}

// ListFormats returns the content of the request which requests for formats
func (repository *EPrintsRepository) ListFormats() ([]byte, error) {
	return repository.request.GetFormats()
}

// ListCollections returns the content of the request which requests the sets
func (repository *EPrintsRepository) ListCollections() ([]byte, error) {
	return repository.request.GetSets()
}

// ListRecords returns the content of the request which requests for records
func (repository *EPrintsRepository) ListRecords(token string) ([]byte, error) {
	repository.request.Clear()
	repository.request.ResumptionToken = token
	return repository.request.GetRecords("oai_dc")
}

// ListIdentifiers returns the content of the request which requests for identifiers
func (repository *EPrintsRepository) ListIdentifiers(token string) ([]byte, error) {
	repository.request.Clear()
	repository.request.ResumptionToken = token
	return repository.request.GetIdentifiers("oai_dc")
}

// ----------------------------------- Parse

// GetIdentification returns the identification details in Wisply format
func (repository *EPrintsRepository) GetIdentification(content []byte) (*wisply.Identificationer, error) {

	var idenfitication wisply.Identificationer

	response, err := repository.request.Parse(content)
	if err != nil {
		return &idenfitication, err
	}

	remoteIdentity := response.Identify

	// check a field which must be
	if remoteIdentity.RepositoryName == "" {
		return &idenfitication, errors.New("There was a problem getting the fields")
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
func (repository *EPrintsRepository) GetFormats(content []byte) ([]wisply.Formater, error) {

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
func (repository *EPrintsRepository) GetCollections(content []byte) ([]wisply.Collectioner, error) {
	var collections []wisply.Collectioner

	response, err := repository.request.Parse(content)
	if err != nil {
		return collections, err
	}

	remoteCollections := response.ListSets.Set

	for _, collection := range remoteCollections {
		collection := &Collection{
			Name: collection.SetName,
			Spec: collection.SetSpec,
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

// GetRecords returns the `records` in Wisply format
func (repository *EPrintsRepository) GetRecords(content []byte) ([]wisply.Recorder, error) {

	var records []wisply.Recorder

	response, err := repository.request.Parse(content)

	// cache the last response
	repository.lastResponse = response

	if err != nil {
		return records, err
	}

	remoteRecords := response.ListRecords.Records

	for _, record := range remoteRecords {

		keys, err := repository.getKeys(record.Metadata.Body)

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

func (repository *EPrintsRepository) getKeys(plainText []byte) (Keys, error) {

	keys := Keys{}

	err := xml.Unmarshal(plainText, &keys)
	if err != nil {
		return keys, err
	}
	return keys, nil
}

// GetIdentifiers returns the `records` in Wisply format
func (repository *EPrintsRepository) GetIdentifiers(content []byte) ([]wisply.Identifier, error) {

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
func (repository *EPrintsRepository) GetNextPage() (string, bool) {
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
func (repository *EPrintsRepository) GetFinishToken() string {
	return repository.finishToken
}

func (repository *EPrintsRepository) prepareFinishToken() {
	resp := repository.lastResponse

	record := resp.ListRecords.Records[len(resp.ListRecords.Records)-1]
	identifier := record.Header.Identifier

	elements := strings.Split(identifier, ":")

	id, _ := strconv.Atoi(elements[2])

	repository.finishToken = "metadataPrefix%3Doai_dc%26offset%3D" + strconv.Itoa(id+1)
}

// IsValidResponse checks if the content is an OAI format
func (repository *EPrintsRepository) IsValidResponse(content []byte) error {
	return repository.request.IsValidResponse(content)
}

// NewEPrints returns a repository of type Eprints
func NewEPrints(rep *repository.Repository) wisply.RepositoryInterface {
	req := &protocol.Request{
		BaseURL: rep.GetLocalRepository().URL,
	}
	return &EPrintsRepository{
		Repository: rep,
		request:    req,
		filter:     newFilter(rep.GetLocalRepository().GetFilter()),
	}
}
