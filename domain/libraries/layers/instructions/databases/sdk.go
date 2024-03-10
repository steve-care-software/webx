package databases

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/reverts"
)

// Service represents a service
type Service interface {
	IsBegin() bool
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
	IsCommit() bool
	Commit() string
	IsCancel() bool
	Cancel() string
	IsRevert() bool
	Revert() reverts.Revert
}
