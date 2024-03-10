package inserts

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithPassword(password string) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Username() string
	Password() string
}
