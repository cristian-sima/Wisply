package harvest

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
	GetID() int64
	SetID(ID int64)
}
