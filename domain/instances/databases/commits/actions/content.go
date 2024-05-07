package actions

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values"

type content struct {
	isDelete bool
	insert   values.Value
}

func createContentWithDelete() Content {
	return createContentInternally(true, nil)
}

func createContentWithInsert(
	insert values.Value,
) Content {
	return createContentInternally(false, insert)
}

func createContentInternally(
	isDelete bool,
	insert values.Value,
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
func (obj *content) Insert() values.Value {
	return obj.insert
}
