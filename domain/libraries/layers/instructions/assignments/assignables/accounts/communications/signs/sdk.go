package signs

// Builder represents a sign builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithAccount(account string) Builder
	Now() (Sign, error)
}

// Sign represenst a sign
type Sign interface {
	Message() string
	Account() string
}
