package remote

// RepositoryInterface ... defines the method to be implemented by a remote repository
type RepositoryInterface interface {
	IsValid() ([]byte, error)
	// Identify()
	// HarvestFormats()
	// HarvestCollections()
	// HarvestRecords()
	// HarvestIdentifiers()
}
