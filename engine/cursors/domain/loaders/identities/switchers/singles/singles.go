package singles

type singles struct {
	list []Single
}

func createSingles(
	list []Single,
) Singles {
	out := singles{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *singles) List() []Single {
	return obj.list
}
