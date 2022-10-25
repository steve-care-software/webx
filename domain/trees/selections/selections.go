package selections

type selections struct {
	treeName string
	list     []Selection
}

func createSelections(
	treeName string,
	list []Selection,
) Selections {
	out := selections{
		treeName: treeName,
		list:     list,
	}

	return &out
}

// TreeName returns the treeName
func (obj *selections) TreeName() string {
	return obj.treeName
}

// List returns the selection list
func (obj *selections) List() []Selection {
	return obj.list
}

// Bytes returns the bytes
func (obj *selections) Bytes() []byte {
	output := []byte{}
	for _, oneSelection := range obj.list {
		output = append(output, oneSelection.Bytes()...)
	}

	return output
}
