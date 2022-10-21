package outputs

type output struct {
	values    map[string]interface{}
	remaining []byte
}

func createOutput(
	values map[string]interface{},
) Output {
	return createOutputInternally(values, nil)
}

func createOutputWithRemaining(
	values map[string]interface{},
	remaining []byte,
) Output {
	return createOutputInternally(values, remaining)
}

func createOutputInternally(
	values map[string]interface{},
	remaining []byte,
) Output {
	out := output{
		values:    values,
		remaining: remaining,
	}

	return &out
}

// Values returns the values
func (obj *output) Values() map[string]interface{} {
	return obj.values
}

// HasRemaining returns true if there is a remaining, false otherwise
func (obj *output) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining, if any
func (obj *output) Remaining() []byte {
	return obj.remaining
}
