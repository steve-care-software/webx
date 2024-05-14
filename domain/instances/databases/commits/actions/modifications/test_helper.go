package modifications

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"

// NewModificationsForTests creates new modifications for tests
func NewModificationsForTests(list []Modification) Modifications {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithInsertForTests creates a new modification with insert for tests
func NewModificationWithInsertForTests(insert []byte) Modification {
	ins, err := NewModificationBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithDeleteForTests creates a new modification with delete for tests
func NewModificationWithDeleteForTests(delete deletes.Delete) Modification {
	ins, err := NewModificationBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
