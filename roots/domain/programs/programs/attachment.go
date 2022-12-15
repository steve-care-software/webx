package programs

import "github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"

type attachment struct {
	value Value
	local uint
}

func createAttachment(
	value Value,
	local uint,
) Attachment {
	out := attachment{
		value: value,
		local: local,
	}

	return &out
}

// Hash returns the hash
func (obj *attachment) Hash() hash.Hash {
	return obj.value.Hash()
}

// Value returns the value
func (obj *attachment) Value() Value {
	return obj.value
}

// Local returns the local
func (obj *attachment) Local() uint {
	return obj.local
}
