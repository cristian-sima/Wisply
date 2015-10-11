package remote

import (
	"github.com/cristian-sima/Wisply/models/repository"
)

// RepositoryInterface ... defines the method to be implemented by a remote repository
type RepositoryInterface interface {
	// Validate()
	// Identify()
	// HarvestFormats()
	// HarvestCollections()
	// HarvestRecords()
	// HarvestIdentifiers()
}

// NewEprints returns a repository of type Eprints
func NewEPrints(rep *repository.Repository) RepositoryInterface {
	return &EPrintsRepository{
		Repository: &Repository{
			repository: rep,
		},
	}
}
