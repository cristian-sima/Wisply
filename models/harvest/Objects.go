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

// Collection ... must be implemented by a repository
type Collection interface {
	GetName() string
	GetSpec() string
}

// RESOURCES

// Record ... must be implemented by a repository
type Record interface {
	GetIdentifier() string
	GetDatestamp() string
	GetKeys() *Keys
	GetID() int64
	SetID(ID int64)
}
