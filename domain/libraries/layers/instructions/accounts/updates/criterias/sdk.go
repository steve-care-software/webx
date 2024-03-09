package criterias

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials string) Builder
	ChangeSigner() Builder
	ChangeEncryptor() Builder
	Now() (Criteria, error)
}

// Criteria represents an update criteria
type Criteria interface {
	ChangeSigner() bool
	ChangeEncryptor() bool
	Credentials() string
}
