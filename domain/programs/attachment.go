package programs

type attachment struct {
	value Value
	local []byte
}

func createAttachment(
	value Value,
	local []byte,
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
func (obj *attachment) Local() []byte {
	return obj.local
}
