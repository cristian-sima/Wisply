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

// RemoteRepository represents a remote repository
type RemoteRepository struct {
	Manager *Process
}

// SetManager sets the manager of a current repository
func (repository *RemoteRepository) SetManager(manager *Process) {
	repository.Manager = manager
}

// Identificationer ... must be implemented by a repository
type Identificationer interface {
	GetName() string
	GetURL() string
	GetGranularity() string
	GetDeletedRecord() string
	GetProtocol() string
	GetEarliestDatestamp() string
	GetAdminEmails() []string
}

// Formater ... must be implemented by a repository
type Formater interface {
	GetPrefix() string
	GetNamespace() string
	GetSchema() string
}

// COLLECTIONS

// Collectioner ... must be implemented by a repository
type Collectioner interface {
	GetName() string
	GetSpec() string
}

// RESOURCES

// Recorder ... must be implemented by a repository
type Recorder interface {
	GetIdentifier() string
	GetDatestamp() string
	GetKeys() *Keys
}

// IDENTIFIERS

// Identifier ... must be implemented by a identifier
type Identifier interface {
	GetIdentifier() string
	GetTimestamp() string
	GetSpec() []string
}
