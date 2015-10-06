package harvest

// IdentificationResulter ... is a result from the remote repository
type IdentificationResulter interface {
	IsOk() bool
	GetData() *Identificationer
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

// FormatResulter ... is a result from the remote repository
type FormatResulter interface {
	IsOk() bool
	GetData() []Formater
}

// Formater ... must be implemented by a repository
type Formater interface {
	GetPrefix() string
	GetNamespace() string
	GetSchema() string
}

// COLLECTIONS

// CollectionResult ... is a result from the remote repository
type CollectionResult interface {
	IsOk() bool
	GetData() []Collection
}

// Collection ... must be implemented by a repository
type Collection interface {
	GetName() string
	GetSpec() string
}

// RESOURCES

// RecordResult ... is a result from the remote repository
type RecordResult interface {
	IsOk() bool
	GetData() []Record
}

// Record ... must be implemented by a repository
type Record interface {
	GetIdentifier() string
	GetDatestamp() string
	GetKeys() *Keys
	GetID() int64
	SetID(ID int64)
}
