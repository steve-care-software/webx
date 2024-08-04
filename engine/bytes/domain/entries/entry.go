package entries

import "github.com/steve-care-software/webx/engine/bytes/domain/delimiters"

type entry struct {
	delimiter delimiters.Delimiter
	bytes     []byte
}

func createEntry(
	delimiter delimiters.Delimiter,
	bytes []byte,
) Entry {
	out := entry{
		delimiter: delimiter,
		bytes:     bytes,
	}

	return &out
}

// Delimiter returns the delimiter
func (obj *entry) Delimiter() delimiters.Delimiter {
	return obj.delimiter
}

// Bytes returns the bytes
func (obj *entry) Bytes() []byte {
	return obj.bytes
}
