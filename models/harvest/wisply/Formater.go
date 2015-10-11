package wisply

// Formater ... defines the methods of the formats the repository
type Formater interface {
	GetPrefix() string
	GetNamespace() string
	GetSchema() string
}
