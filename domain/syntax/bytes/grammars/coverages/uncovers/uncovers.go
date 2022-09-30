package uncovers

type uncovers struct {
	list []Uncover
}

func createUncovers(
	list []Uncover,
) Uncovers {
	out := uncovers{
		list: list,
	}

	return &out
}

// List returns the uncovers
func (obj *uncovers) List() []Uncover {
	return obj.list
}
