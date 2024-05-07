package actions

import (
	"github.com/steve-care-software/datastencil/domain/instances"
)

type content struct {
	isDelete bool
	insert   instances.Instance
}

func createContentWithDelete() Content {
	return createContentInternally(true, nil)
}

func createContentWithInsert(
	insert instances.Instance,
) Content {
	return createContentInternally(false, insert)
}

func createContentInternally(
	isDelete bool,
	insert instances.Instance,
) Content {
	out := content{
		isDelete: isDelete,
		insert:   insert,
	}

	return &out
}

// IsDelete returns true if delete, false otherwise
func (obj *content) IsDelete() bool {
	return obj.isDelete
}

// IsInsert returns true if insert, false otherwise
func (obj *content) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *content) Insert() instances.Instance {
	return obj.insert
}
