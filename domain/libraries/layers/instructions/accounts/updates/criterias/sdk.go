package criterias

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password string) Builder
	ChangeSigner() Builder
	ChangeEncryptor() Builder
	Now() (Criteria, error)
}

// Criteria represents an update criteria
type Criteria interface {
	ChangeSigner() bool
	ChangeEncryptor() bool
	HasUsername() bool
	Username() string
	HasPassword() bool
	Password() string
}
