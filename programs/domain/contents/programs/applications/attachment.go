package applications

import "github.com/steve-care-software/webx/domain/databases/entities"

type attachment struct {
	value entities.Identifier
	local uint
}

func createAttachment(
	value entities.Identifier,
	local uint,
) Attachment {
	out := attachment{
		value: value,
		local: local,
	}

	return &out
}

// Value returns the value
func (obj *attachment) Value() entities.Identifier {
	return obj.value
}

// Local returns the local
func (obj *attachment) Local() uint {
	return obj.local
}
