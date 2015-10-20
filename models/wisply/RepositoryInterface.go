package wisply

// RepositoryInterface ... defines the method to be implemented by a remote repository
type RepositoryInterface interface {
	Test() ([]byte, error)
	IsValidResponse(content []byte) error

	Identify() ([]byte, error)
	GetIdentification(content []byte) (*Identificationer, error)

	ListFormats() ([]byte, error)
	GetFormats(content []byte) ([]Formater, error)

	ListCollections(token string) ([]byte, error)
	GetCollections(content []byte) ([]Collectioner, error)

	ListRecords(token string) ([]byte, error)
	GetRecords(content []byte) ([]Recorder, error)

	ListIdentifiers(token string) ([]byte, error)
	GetIdentifiers(content []byte) ([]Identifier, error)

	// GetResumptionToken() string

	GetNextPage() (string, bool)

	GetFinishToken() string
}
