package criterias

import "time"

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithAmount(amount uint) Builder
	WithFrom(from time.Time) Builder
	WithTo(to time.Time) Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Amount() uint
	HasFrom() bool
	From() time.Time
	HasTo() bool
	To() time.Time
}
