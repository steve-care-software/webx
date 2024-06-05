package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
)

// NewDatabaseForTests creates a database for tests
func NewDatabaseForTests(commit commits.Commit, head heads.Head) Database {
	builder := NewBuilder().Create().WithCommit(commit).WithHead(head)
	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
