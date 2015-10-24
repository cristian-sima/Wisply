package data

// Recorder ... must be implemented by a repository
type Recorder interface {
	GetIdentifier() string
	GetTimestamp() string
	GetKeys() Keys
}

// Keys encapsulate all the dublin core keys
type Keys interface {
	GetTitles() []string
	GetCreators() []string
	GetSubjects() []string
	GetDescriptions() []string
	GetPublishers() []string
	GetContributors() []string
	GetDates() []string
	GetTypes() []string
	GetFormats() []string
	GetIdentifiers() []string
	GetSources() []string
	GetLanguages() []string
	GetRelations() []string
	GetCoverages() []string
	GetRights() []string
}
