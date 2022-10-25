package selections

import "github.com/steve-care-software/webx/domain/trees"

type element struct {
	value               trees.Element
	includeChannelBytes bool
}

func createElement(
	value trees.Element,
	includeChannelBytes bool,
) Element {
	out := element{
		value:               value,
		includeChannelBytes: includeChannelBytes,
	}

	return &out
}

// Value returns the value
func (obj *element) Value() trees.Element {
	return obj.value
}

// IncludeChannelBytes returns true if the channel bytes are included, false otherwise
func (obj *element) IncludeChannelBytes() bool {
	return obj.includeChannelBytes
}

// Bytes returns the bytes
func (obj *element) Bytes() []byte {
	return obj.value.Bytes(obj.includeChannelBytes)
}
