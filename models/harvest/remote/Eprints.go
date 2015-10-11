package remote

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/harvest/remote/protocols/oai"
	"github.com/cristian-sima/Wisply/models/harvest/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
)

// EPrintsRepository is an Eprints remote repository
type EPrintsRepository struct {
	*Repository
	request *oai.Request
}

// Test calls the request validate method and returns its result
func (repository *EPrintsRepository) Test() ([]byte, error) {
	return repository.request.Identify()
}

// IsValidResponse checks if the content is a valid OAI reponse type
func (repository *EPrintsRepository) IsValidResponse(content []byte) error {
	return repository.request.IsValidResponse(content)
}

// Identify returns the identification details
func (repository *EPrintsRepository) Identify() ([]byte, error) {
	return repository.request.Identify()
}

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

	idenfitication = &OAIIdentification{
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

// NewEPrints returns a repository of type Eprints
func NewEPrints(rep *repository.Repository) RepositoryInterface {
	req := &oai.Request{
		BaseURL: rep.URL,
	}
	return &EPrintsRepository{
		Repository: &Repository{
			repository: rep,
		},
		request: req,
	}
}