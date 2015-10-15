package remote

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/harvest/remote/oai"
	harvestRepository "github.com/cristian-sima/Wisply/models/harvest/remote/repository"
	"github.com/cristian-sima/Wisply/models/repository"
)

// New creates a new remote server based on the information provided by the local one
// It is a `factory` pattern
func New(localRepository *repository.Repository) (RepositoryInterface, error) {
	var rem RepositoryInterface

	basic := &harvestRepository.Repository{}

	basic.SetLocalRepository(localRepository)

	switch localRepository.Category {
	case "EPrints":
		{
			rem = oai.NewEPrints(basic)
		}
	default:
		return rem, errors.New("There is no such server like " + localRepository.Category)
	}
	return rem, nil
}
