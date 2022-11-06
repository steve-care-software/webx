package entities

type identifiers struct {
	list []Identifier
}

func createIdentifiers(
	list []Identifier,
) Identifiers {
	out := identifiers{
		list: list,
	}

	return &out
}

// List returns the identifiers
func (obj *identifiers) List() []Identifier {
	return obj.list
}
