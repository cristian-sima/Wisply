package oai

import (
	wisply "github.com/cristian-sima/Wisply/models/wisply/data"
)

// Identification represents the identification for a repository
type Identification struct {
	Name              string
	URL               string
	Granularity       string
	RecordPolicy      string
	Protocol          string
	EarliestDatestamp string
	AdminEmails       []string
}

// GetURL returns the URL
func (repository *Identification) GetURL() string {
	return repository.URL
}

// GetName returns the name
func (repository *Identification) GetName() string {
	return repository.Name
}

// GetDeletedRecord returns the protocol for record deletion
func (repository *Identification) GetDeletedRecord() string {
	return repository.RecordPolicy
}

// GetGranularity returns the granularity
func (repository *Identification) GetGranularity() string {
	return repository.Granularity
}

// GetProtocol returns teh protocol of repository
func (repository *Identification) GetProtocol() string {
	return repository.Protocol
}

// GetEarliestDatestamp retuns the last timestamp of the repository
func (repository *Identification) GetEarliestDatestamp() string {
	return repository.EarliestDatestamp
}

// GetAdminEmails returns the list of admin emails for the repository
func (repository *Identification) GetAdminEmails() []string {
	return repository.AdminEmails
}

// FORMATS

// Format represents the identification for a  repository
type Format struct {
	Prefix    string
	Schema    string
	Namespace string
}

// GetSchema returns the schema
func (repository *Format) GetSchema() string {
	return repository.Schema
}

// GetPrefix returns the prefix
func (repository *Format) GetPrefix() string {
	return repository.Prefix
}

// GetNamespace returns the namespace
func (repository *Format) GetNamespace() string {
	return repository.Namespace
}

// COLLECTIONS

// Collection represents an  collection
type Collection struct {
	Name string
	Spec string
}

// GetName returns the name of the collection
func (repository *Collection) GetName() string {
	return repository.Name
}

// GetSpec returns the id of the collection
func (repository *Collection) GetSpec() string {
	return repository.Spec
}

// RESOURCES

// Keys encapsulate all the dublin core keys
type Keys struct {
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
func (keys Keys) GetTitles() []string {
	return keys.Titles
}

// GetCreators returns the Creators
func (keys Keys) GetCreators() []string {
	return keys.Creators
}

// GetSubjects returns the Subjects
func (keys Keys) GetSubjects() []string {
	return keys.Subjects
}

// GetDescriptions returns the Descriptions
func (keys Keys) GetDescriptions() []string {
	return keys.Descriptions
}

// GetPublishers returns the Publishers
func (keys Keys) GetPublishers() []string {
	return keys.Publishers
}

// GetContributors returns the Contributors
func (keys Keys) GetContributors() []string {
	return keys.Contributors
}

// GetDates returns the Dates
func (keys Keys) GetDates() []string {
	return keys.Dates
}

// GetTypes returns the Types
func (keys Keys) GetTypes() []string {
	return keys.Types
}

// GetFormats returns the Formats
func (keys Keys) GetFormats() []string {
	return keys.Formats
}

// GetIdentifiers returns the Identifiers
func (keys Keys) GetIdentifiers() []string {
	return keys.Identifiers
}

// GetSources returns the Sources
func (keys Keys) GetSources() []string {
	return keys.Sources
}

// GetLanguages returns the Languages
func (keys Keys) GetLanguages() []string {
	return keys.Languages
}

// GetRelations returns the Relations
func (keys Keys) GetRelations() []string {
	return keys.Relations
}

// GetCoverages returns the Coverages
func (keys Keys) GetCoverages() []string {
	return keys.Coverages
}

// GetRights returns the Rights
func (keys Keys) GetRights() []string {
	return keys.Rights
}

// Record represents an  record
type Record struct {
	Datestamp  string
	Identifier string
	Keys       Keys
}

// GetTimestamp returns the datastamp of the record
func (record Record) GetTimestamp() string {
	return record.Datestamp
}

// GetIdentifier returns the identifier of the record
func (record Record) GetIdentifier() string {
	return record.Identifier
}

// GetKeys returns the keys of the record
func (record Record) GetKeys() wisply.Keys {
	return record.Keys
}

// Identifier represents an  identifier
type Identifier struct {
	Datestamp  string
	Identifier string
	Spec       []string
}

// GetTimestamp returns the datastamp of the record
func (record *Identifier) GetTimestamp() string {
	return record.Datestamp
}

// GetIdentifier returns the identifier of the record
func (record *Identifier) GetIdentifier() string {
	return record.Identifier
}

// GetSpec returns the specs for the identifier
func (record *Identifier) GetSpec() []string {
	return record.Spec
}
