package harvest

// OAIIdentification represents the identification for a OAI repository
type OAIIdentification struct {
	name              string
	url               string
	granularity       string
	recordPolicy      string
	protocol          string
	earliestDatestamp string
	adminEmails       []string
}

// GetURL returns the URL
func (repository *OAIIdentification) GetURL() string {
	return repository.url
}

// GetName returns the name
func (repository *OAIIdentification) GetName() string {
	return repository.name
}

// GetDeletedRecord returns the protocol for record deletion
func (repository *OAIIdentification) GetDeletedRecord() string {
	return repository.recordPolicy
}

// GetGranularity returns the granularity
func (repository *OAIIdentification) GetGranularity() string {
	return repository.recordPolicy
}

// GetProtocol returns teh protocol of repository
func (repository *OAIIdentification) GetProtocol() string {
	return repository.protocol
}

// GetEarliestDatestamp retuns the last timestamp of the repository
func (repository *OAIIdentification) GetEarliestDatestamp() string {
	return repository.earliestDatestamp
}

// GetAdminEmails returns the list of admin emails for the repository
func (repository *OAIIdentification) GetAdminEmails() []string {
	return repository.adminEmails
}

// OAIIdentificationResult is a result of identification for OAI protocol
type OAIIdentificationResult struct {
	isOk bool
	data Identificationer
}

// GetState returns if the result has succeeded
func (identification *OAIIdentificationResult) IsOk() bool {
	return identification.isOk
}

// GetData returns the data
func (identification *OAIIdentificationResult) GetData() *Identificationer {
	return &identification.data
}
