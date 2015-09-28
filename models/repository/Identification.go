package repository

// Identification represents the identification for a repository
type Identification struct {
	ID                int
	Repository        int
	Name              string
	URL               string
	Granularity       string
	RecordPolicy      string
	Protocol          string
	EarliestDatestamp string
	AdminEmails       []string
}
