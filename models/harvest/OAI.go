package harvest

// OAIIdentification represents the identification for a OAI repository
type OAIIdentification struct {
	Name              string
	URL               string
	Granularity       string
	RecordPolicy      string
	Protocol          string
	EarliestDatestamp string
	AdminEmails       []string
}

// GetURL returns the URL
func (repository *OAIIdentification) GetURL() string {
	return repository.URL
}

// GetName returns the name
func (repository *OAIIdentification) GetName() string {
	return repository.Name
}

// GetDeletedRecord returns the protocol for record deletion
func (repository *OAIIdentification) GetDeletedRecord() string {
	return repository.RecordPolicy
}

// GetGranularity returns the granularity
func (repository *OAIIdentification) GetGranularity() string {
	return repository.Granularity
}

// GetProtocol returns teh protocol of repository
func (repository *OAIIdentification) GetProtocol() string {
	return repository.Protocol
}

// GetEarliestDatestamp retuns the last timestamp of the repository
func (repository *OAIIdentification) GetEarliestDatestamp() string {
	return repository.EarliestDatestamp
}

// GetAdminEmails returns the list of admin emails for the repository
func (repository *OAIIdentification) GetAdminEmails() []string {
	return repository.AdminEmails
}

// OAIIdentificationResult is a result of identification for OAI protocol
type OAIIdentificationResult struct {
	isOk bool
	data Identificationer
}

// IsOk returns if the result has succeeded
func (identification *OAIIdentificationResult) IsOk() bool {
	return identification.isOk
}

// GetData returns the data
func (identification *OAIIdentificationResult) GetData() *Identificationer {
	return &identification.data
}
