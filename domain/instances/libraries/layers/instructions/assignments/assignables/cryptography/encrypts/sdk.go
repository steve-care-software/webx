package encrypts

// Builder represents an encrypt builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithPassword(password string) Builder
	Now() (Encrypt, error)
}

// Encrypt represents an encrypt
type Encrypt interface {
	Message() string
	Password() string
}
