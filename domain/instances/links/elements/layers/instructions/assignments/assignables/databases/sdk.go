package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/retrieves"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/transforms"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/values"
)

// Database represents a database assignable
type Database interface {
	IsDatabase() bool
	Database() databases.Database
	IsCommit() bool
	Commit() commits.Commit
	IsAction() bool
	Action() actions.Action
	IsValue() bool
	Value() values.Value
	IsTransform() bool
	Transform() transforms.Transform
	IsRetrieve() bool
	Retrieve() retrieves.Retrieve
}
