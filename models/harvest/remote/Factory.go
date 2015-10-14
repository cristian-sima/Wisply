package remote

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// New creates a new remote server based on the information provided by the local one
// It is a `factory` pattern
func New(localRepository *repository.Repository) (wisply.RepositoryInterface, error) {
	var rem wisply.RepositoryInterface

	switch localRepository.Category {
	case "EPrints":
		{
			rem = NewEPrints(localRepository)
		}
	default:
		return rem, errors.New("There is no such server like " + localRepository.Category)
	}
	return rem, nil
}
