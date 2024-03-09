package votes

// Builder represents a vote builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithRing(ring string) Builder
	WithAccount(account string) Builder
	Now() (Vote, error)
}

// Vote represents a vote
type Vote interface {
	Message() string
	Ring() string
	Account() string
}
