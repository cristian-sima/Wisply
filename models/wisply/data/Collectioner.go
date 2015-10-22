package data

// Collectioner ... must be implemented by a repository
type Collectioner interface {
	GetName() string
	GetSpec() string
	GetPath() string
}
