package applications

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type attachment struct {
	value hash.Hash
	local uint
}

func createAttachment(
	value hash.Hash,
	local uint,
) Attachment {
	out := attachment{
		value: value,
		local: local,
	}

	return &out
}

// Value returns the value
func (obj *attachment) Value() hash.Hash {
	return obj.value
}

// Local returns the local
func (obj *attachment) Local() uint {
	return obj.local
}
