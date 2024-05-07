package databases

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits"

// NewDatabaseForTests creates a database for tests
func NewDatabaseForTests(path []string, description string, head commits.Commit, isActive bool) Database {
	builder := NewBuilder().Create().WithPath(path).WithDescription(description).WithHead(head)
	if isActive {
		builder.IsActive()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
