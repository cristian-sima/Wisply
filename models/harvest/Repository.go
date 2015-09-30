package harvest

// RemoteRepositoryInterface ... defines the method to be implemented by a standard (remote repository)
type RemoteRepositoryInterface interface {
	Validate()
	HarvestIdentification()
	HarvestFormats()
	HarvestCollections()
	HarvestRecords()
	SetManager(manager *Process)
}

// Message encapsulates the message to communicate with controller
type Message struct {
	Name       string
	Content    string
	Value      interface{}
	Repository int
}

// RemoteRepository represents a remote repository
type RemoteRepository struct {
	Manager *Process
}

// SetManager sets the manager of a current repository
func (repository *RemoteRepository) SetManager(manager *Process) {
	repository.Manager = manager
}
