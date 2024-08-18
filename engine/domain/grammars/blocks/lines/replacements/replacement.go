package replacements

type replacement struct {
	origin string
	target string
}

func createReplacement(
	origin string,
	target string,
) Replacement {
	out := replacement{
		origin: origin,
		target: target,
	}

	return &out
}

// Origin returns the origin token
func (obj *replacement) Origin() string {
	return obj.origin
}

// Target returns the target ast
func (obj *replacement) Target() string {
	return obj.target
}
