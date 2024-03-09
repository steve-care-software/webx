package updates

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates/criterias"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates/executions"
)

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithCriteria(criteria criterias.Criteria) Builder
	WithExecution(execution executions.Execution) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	IsCriteria() bool
	Criteria() criterias.Criteria
	IsExecution() bool
	Execution() executions.Execution
}
