package entries

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"

type entry struct {
	pointer pointers.Pointer
	bytes   []byte
}

func createEntry(
	pointer pointers.Pointer,
	bytes []byte,
) Entry {
	out := entry{
		pointer: pointer,
		bytes:   bytes,
	}

	return &out
}

// Pointer returns the pointer
func (obj *entry) Pointer() pointers.Pointer {
	return obj.pointer
}

// Bytes returns the bytes
func (obj *entry) Bytes() []byte {
	return obj.bytes
}
