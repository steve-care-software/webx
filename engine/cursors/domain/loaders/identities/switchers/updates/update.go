package updates

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"

type update struct {
	single singles.Single
	bytes  []byte
}

func createUpdate(
	single singles.Single,
	bytes []byte,
) Update {
	out := update{
		single: single,
		bytes:  bytes,
	}

	return &out
}

// Single returns the single identity instance
func (obj *update) Single() singles.Single {
	return obj.single
}

// Bytes returns the bytes
func (obj *update) Bytes() []byte {
	return obj.bytes
}
