package databases

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits"

// NewDatabaseForTests creates a database for tests
func NewDatabaseForTests(commit commits.Commit) Database {
	builder := NewBuilder().Create().WithCommit(commit)
	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
