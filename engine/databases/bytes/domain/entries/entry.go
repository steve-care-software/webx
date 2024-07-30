package entries

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"

type entry struct {
	keyname   string
	delimiter delimiters.Delimiter
	bytes     []byte
}

func createEntry(
	keyname string,
	delimiter delimiters.Delimiter,
	bytes []byte,
) Entry {
	out := entry{
		keyname:   keyname,
		delimiter: delimiter,
		bytes:     bytes,
	}

	return &out
}

// Keyname returns the keyname
func (obj *entry) Keyname() string {
	return obj.keyname
}

// Delimiter returns the delimiter
func (obj *entry) Delimiter() delimiters.Delimiter {
	return obj.delimiter
}

// Bytes returns the bytes
func (obj *entry) Bytes() []byte {
	return obj.bytes
}
