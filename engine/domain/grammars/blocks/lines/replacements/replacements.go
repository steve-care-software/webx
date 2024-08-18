package replacements

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

// List returns the list of replacement
func (obj *replacements) List() []Replacement {
	return obj.list
}
