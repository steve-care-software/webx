package assignments

import "github.com/steve-care-software/webx/domain/databases/entities"

type assignment struct {
	index uint
	value entities.Identifier
}

func createAssignment(
	index uint,
	value entities.Identifier,
) Assignment {
	out := assignment{
		index: index,
		value: value,
	}

	return &out
}

// Index returns the index
func (obj *assignment) Index() uint {
	return obj.index
}

// Value returns the value
func (obj *assignment) Value() entities.Identifier {
	return obj.value
}
