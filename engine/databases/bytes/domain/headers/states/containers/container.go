package containers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers/pointers"

type container struct {
	keyname  string
	pointers pointers.Pointers
}

func createContainer(
	keyname string,
	pointers pointers.Pointers,
) Container {
	out := container{
		keyname:  keyname,
		pointers: pointers,
	}

	return &out
}

// Keyname returns the keyname
func (obj *container) Keyname() string {
	return obj.keyname
}

// Pointers returns the pointers
func (obj *container) Pointers() pointers.Pointers {
	return obj.pointers
}
