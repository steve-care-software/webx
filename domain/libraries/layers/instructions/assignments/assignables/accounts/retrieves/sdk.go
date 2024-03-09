package retrieves

// Builder represents a retrieve builder
type Builder interface {
	Create() Builder
	WithPassword(password string) Builder
	WithCredentials(credentials string) Builder
	Now() (Retrieve, error)
}

// Retrieve represents a retrieve
type Retrieve interface {
	Password() string
	Credentials() string
}
