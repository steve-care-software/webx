package updates

import "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates/criterias"

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithCredentials(credentials string) Builder
	WithCriteria(criteria criterias.Criteria) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Credentials() string
	Criteria() criterias.Criteria
}
