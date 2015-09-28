package harvest

// RemoteRepositoryInterface ... defines the method to be implemented by a standard (remote repository)
type RemoteRepositoryInterface interface {
	Start()
	SetManager(manager *Manager)
}

// Controller represents a controller
type Controller interface {
	Notify(*Message)
}

// Message encapsulates the message to communicate with controller
type Message struct {
	Name       string
	Content    string
	Value      string
	Repository int
}

// RemoteRepository represents a remote repository
type RemoteRepository struct {
	Manager *Manager
}

// SetManager sets the manager of a current repository
func (repository *RemoteRepository) SetManager(manager *Manager) {
	repository.Manager = manager
}
