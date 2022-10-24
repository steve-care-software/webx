package programs

type attachment struct {
	value Value
	local string
}

func createAttachment(
	value Value,
	local string,
) Attachment {
	out := attachment{
		value: value,
		local: local,
	}

	return &out
}

// Value returns the value
func (obj *attachment) Value() Value {
	return obj.value
}

// Local returns the local
func (obj *attachment) Local() string {
	return obj.local
}
