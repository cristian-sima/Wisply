package remote

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// RepositoryInterface ... defines the method to be implemented by a remote repository
type RepositoryInterface interface {
	Test() ([]byte, error)
	IsValidResponse(content []byte) error
	Identify() ([]byte, error)
	GetIdentification(content []byte) (*wisply.Identificationer, error)
}
