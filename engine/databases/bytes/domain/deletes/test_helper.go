package deletes

import "github.com/steve-care-software/webx/engine/states/domain/databases/pointers"

// NewDeletesForTests creates a new deletes for tests
func NewDeletesForTests(list []Delete) Deletes {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDeleteForTests creates a new delete for tests
func NewDeleteForTests(keyname string, pointer pointers.Pointer) Delete {
	ins, err := NewDeleteBuilder().Create().WithKeyname(keyname).WithPointer(pointer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
