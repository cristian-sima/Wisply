package harvest

// IdentificationResultInterface ... is a result from the remote repository
type IdentificationResulter interface {
	IsOk() bool
	GetData() *Identificationer
}

// IdentificationInterface ... must be implemented by a repository
type Identificationer interface {
	GetName() string
	GetURL() string
	GetGranularity() string
	GetDeletedRecord() string
	GetProtocol() string
	GetEarliestDatestamp() string
	GetAdminEmails() []string
}
