package harvest

import "github.com/cristian-sima/Wisply/models/harvest/remote"

// Identificationer ... defines the methods of a identification object
type Identificationer interface {
	GetName() string
	GetURL() string
	GetGranularity() string
	GetDeletedRecord() string
	GetProtocol() string
	GetEarliestDatestamp() string
	GetAdminEmails() []string
}

// Formater ... defines the methods of the formats the repository
type Formater interface {
	GetPrefix() string
	GetNamespace() string
	GetSchema() string
}

// Collectioner ... must be implemented by a repository
type Collectioner interface {
	GetName() string
	GetSpec() string
}

// Recorder ... must be implemented by a repository
type Recorder interface {
	GetIdentifier() string
	GetDatestamp() string
	GetKeys() *remote.Keys
}

// Identifier ... must be implemented by a identifier
type Identifier interface {
	GetIdentifier() string
	GetTimestamp() string
	GetSpec() []string
}
