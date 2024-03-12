package updates

import "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts/updates/criterias"

// NewUpdateForTests creates a new update for tests
func NewUpdateForTests(credentials string, criteria criterias.Criteria) Update {
	ins, err := NewBuilder().Create().WithCredentials(credentials).WithCriteria(criteria).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
