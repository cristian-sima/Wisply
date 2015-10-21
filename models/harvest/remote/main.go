// Package remote encapsulates the functionality for a remote repository
// A remote repository is a web server which delivers metadata.
// Wispy collects the metadata which is afterwards processed
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
		message := "There is no such server like " + localRepository.Category
		return rem, errors.New(message)
	}
	return rem, nil
}
