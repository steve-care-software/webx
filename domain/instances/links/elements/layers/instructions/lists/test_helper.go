package lists

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/inserts"
)

// NewListWithInsertForTests creates a new list with insert for tests
func NewListWithInsertForTests(insert inserts.Insert) List {
	ins, err := NewBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewListWithDeleteForTests creates a new list with delete for tests
func NewListWithDeleteForTests(delete deletes.Delete) List {
	ins, err := NewBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
