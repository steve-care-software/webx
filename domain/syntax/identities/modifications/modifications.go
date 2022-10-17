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
