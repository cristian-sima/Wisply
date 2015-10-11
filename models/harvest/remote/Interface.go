package remote

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// RepositoryInterface ... defines the method to be implemented by a remote repository
type RepositoryInterface interface {
	Test() ([]byte, error)
	IsValidResponse(content []byte) error

	Identify() ([]byte, error)
	GetIdentification(content []byte) (*wisply.Identificationer, error)

	ListFormats() ([]byte, error)
	GetFormats(content []byte) ([]wisply.Formater, error)

	ListCollections() ([]byte, error)
	GetCollections(content []byte) ([]wisply.Collectioner, error)

	ListRecords() ([]byte, error)
	GetRecords(content []byte) ([]wisply.Recorder, error)

	ListIdentifiers() ([]byte, error)
	GetIdentifiers(content []byte) ([]wisply.Identifier, error)
}
