package deletes

import "github.com/steve-care-software/webx/engine/states/domain/databases/pointers"

type delete struct {
	keyname string
	pointer pointers.Pointer
}

func createDelete(
	keyname string,
	pointer pointers.Pointer,
) Delete {
	out := delete{
		keyname: keyname,
		pointer: pointer,
	}

	return &out
}

// Keyname returns the keyname
func (obj *delete) Keyname() string {
	return obj.keyname
}

// Pointer returns the pointer
func (obj *delete) Pointer() pointers.Pointer {
	return obj.pointer
}
