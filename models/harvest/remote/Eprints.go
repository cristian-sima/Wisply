package remote

import (
	"github.com/cristian-sima/Wisply/models/harvest/remote/protocols/oai"
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
