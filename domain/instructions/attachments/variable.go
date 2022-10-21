package attachments

type variable struct {
	current string
	target  string
}

func createVariable(
	current string,
	target string,
) Variable {
	out := variable{
		current: current,
		target:  target,
	}

	return &out
}

// Current returns the current variable
func (obj *variable) Current() string {
	return obj.current
}

// Target returns the target variable
func (obj *variable) Target() string {
	return obj.target
}
