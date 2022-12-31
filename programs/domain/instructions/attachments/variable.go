package attachments

type variable struct {
	current []byte
	target  []byte
}

func createVariable(
	current []byte,
	target []byte,
) Variable {
	out := variable{
		current: current,
		target:  target,
	}

	return &out
}

// Current returns the current variable
func (obj *variable) Current() []byte {
	return obj.current
}

// Target returns the target variable
func (obj *variable) Target() []byte {
	return obj.target
}
