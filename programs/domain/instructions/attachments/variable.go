package attachments

type variable struct {
	current []byte
	target  uint
}

func createVariable(
	current []byte,
	target uint,
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
func (obj *variable) Target() uint {
	return obj.target
}
