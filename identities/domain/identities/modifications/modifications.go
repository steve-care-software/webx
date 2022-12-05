package modifications

type modifications struct {
	list []Modification
}

func createModifications(
	list []Modification,
) Modifications {
	out := modifications{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *modifications) List() []Modification {
	return obj.list
}

// First returns the first modification instance
func (obj *modifications) First() Modification {
	return obj.list[:1][0]
}
