package oai

import (
	"github.com/cristian-sima/Wisply/models/harvest/remote/repository"
	wisply "github.com/cristian-sima/Wisply/models/wisply/data"
)

// EPrintsRepository is an Eprints remote repository
type EPrintsRepository struct {
	*Repository
}

// NewEPrints returns a repository of type Eprints
func NewEPrints(rep *repository.Repository) wisply.RepositoryInterface {
	return &EPrintsRepository{
		Repository: newRepository(rep),
	}
}
