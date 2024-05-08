package actions

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"

type content struct {
	isDelete bool
	insert   modifications.Modifications
}

func createContentWithDelete() Content {
	return createContentInternally(true, nil)
}

func createContentWithInsert(
	insert modifications.Modifications,
) Content {
	return createContentInternally(false, insert)
}

func createContentInternally(
	isDelete bool,
	insert modifications.Modifications,
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
func (obj *content) Insert() modifications.Modifications {
	return obj.insert
}
