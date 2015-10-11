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

// IsValid calls the request validate method and returns its result
func (repository *EPrintsRepository) IsValid() ([]byte, error) {
	return repository.request.Verify()
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
