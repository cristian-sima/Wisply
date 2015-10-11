package remote

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

// FORMATS

// OAIFormat represents the identification for a OAI repository
type OAIFormat struct {
	Prefix    string
	Schema    string
	Namespace string
}

// GetSchema returns the schema
func (repository *OAIFormat) GetSchema() string {
	return repository.Schema
}

// GetPrefix returns the prefix
func (repository *OAIFormat) GetPrefix() string {
	return repository.Prefix
}

// GetNamespace returns the namespace
func (repository *OAIFormat) GetNamespace() string {
	return repository.Namespace
}

// COLLECTIONS

// OAICollection represents an OAI collection
type OAICollection struct {
	Name string
	Spec string
}

// GetName returns the name of the collection
func (repository *OAICollection) GetName() string {
	return repository.Name
}

// GetSpec returns the id of the collection
func (repository *OAICollection) GetSpec() string {
	return repository.Spec
}

// RESOURCES

// OAIKeys encapsulate all the dublin core keys
type OAIKeys struct {
	Titles       []string `xml:"title"`
	Creators     []string `xml:"creator"`
	Subjects     []string `xml:"subject"`
	Descriptions []string `xml:"description"`
	Publishers   []string `xml:"publisher"`
	Contributors []string `xml:"contributor"`
	Dates        []string `xml:"date"`
	Types        []string `xml:"type"`
	Formats      []string `xml:"format"`
	Identifiers  []string `xml:"identifier"`
	Sources      []string `xml:"source"`
	Languages    []string `xml:"language"`
	Relations    []string `xml:"relation"`
	Coverages    []string `xml:"coverage"`
	Rights       []string `xml:"rights"`
}

// GetTitles returns the Titles
func (keys *OAIKeys) GetTitles() []string {
	return keys.Titles
}

// GetCreators returns the Creators
func (keys *OAIKeys) GetCreators() []string {
	return keys.Creators
}

// GetSubjects returns the Subjects
func (keys *OAIKeys) GetSubjects() []string {
	return keys.Subjects
}

// GetDescriptions returns the Descriptions
func (keys *OAIKeys) GetDescriptions() []string {
	return keys.Descriptions
}

// GetPublishers returns the Publishers
func (keys *OAIKeys) GetPublishers() []string {
	return keys.Publishers
}

// GetContributors returns the Contributors
func (keys *OAIKeys) GetContributors() []string {
	return keys.Contributors
}

// GetDates returns the Dates
func (keys *OAIKeys) GetDates() []string {
	return keys.Dates
}

// GetTypes returns the Types
func (keys *OAIKeys) GetTypes() []string {
	return keys.Types
}

// GetFormats returns the Formats
func (keys *OAIKeys) GetFormats() []string {
	return keys.Formats
}

// GetIdentifiers returns the Identifiers
func (keys *OAIKeys) GetIdentifiers() []string {
	return keys.Identifiers
}

// GetSources returns the Sources
func (keys *OAIKeys) GetSources() []string {
	return keys.Sources
}

// GetLanguages returns the Languages
func (keys *OAIKeys) GetLanguages() []string {
	return keys.Languages
}

// GetRelations returns the Relations
func (keys *OAIKeys) GetRelations() []string {
	return keys.Relations
}

// GetCoverages returns the Coverages
func (keys *OAIKeys) GetCoverages() []string {
	return keys.Coverages
}

// GetRights returns the Rights
func (keys *OAIKeys) GetRights() []string {
	return keys.Rights
}

// OAIRecord represents an OAI record
type OAIRecord struct {
	Datestamp  string
	Identifier string
	Keys       OAIKeys
}

// GetTimestamp returns the datastamp of the record
func (record OAIRecord) GetTimestamp() string {
	return record.Datestamp
}

// GetIdentifier returns the identifier of the record
func (record OAIRecord) GetIdentifier() string {
	return record.Identifier
}

// GetKeys returns the keys of the record
func (record *OAIRecord) GetKeys() OAIKeys {
	return record.Keys
}

// OAIIdentifier represents an OAI identifier
type OAIIdentifier struct {
	Datestamp  string
	Identifier string
	Spec       []string
}

// GetTimestamp returns the datastamp of the record
func (record *OAIIdentifier) GetTimestamp() string {
	return record.Datestamp
}

// GetIdentifier returns the identifier of the record
func (record *OAIIdentifier) GetIdentifier() string {
	return record.Identifier
}

// GetSpec returns the specs for the identifier
func (record *OAIIdentifier) GetSpec() []string {
	return record.Spec
}
