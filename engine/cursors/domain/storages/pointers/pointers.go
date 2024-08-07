package pointers

type pointers struct {
	list []Pointer
}

func createPointers(
	list []Pointer,
) Pointers {
	out := pointers{
		list: list,
	}

	return &out
}

// List returns the list of pointers
func (obj *pointers) List() []Pointer {
	return obj.list
}
