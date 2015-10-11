package wisply

// Identifier ... must be implemented by a identifier
type Identifier interface {
	GetIdentifier() string
	GetTimestamp() string
	GetSpec() []string
}
