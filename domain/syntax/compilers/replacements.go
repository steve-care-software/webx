package compilers

type replacements struct {
	list []Replacement
}

func createReplacements(
	list []Replacement,
) Replacements {
	out := replacements{
		list: list,
	}

	return &out
}

// List returns the replacements
func (obj *replacements) List() []Replacement {
	return obj.list
}
