package remote

// RepositoryInterface ... defines the method to be implemented by a remote repository
type RepositoryInterface interface {
	Test() ([]byte, error)
	IsValidResponse(content []byte) error
	// Identify()
	// HarvestFormats()
	// HarvestCollections()
	// HarvestRecords()
	// HarvestIdentifiers()
}
