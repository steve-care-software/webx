package credentials

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the credentials builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password []byte) Builder
	Now() (Credentials, error)
}

// Credentials represents credentials
type Credentials interface {
	Username() string
	Password() []byte
}
