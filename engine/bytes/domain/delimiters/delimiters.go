package delimiters

type delimiters struct {
	list []Delimiter
}

func createDelimiters(
	list []Delimiter,
) Delimiters {
	out := delimiters{
		list: list,
	}

	return &out
}

// List returns the list of delimiters
func (obj *delimiters) List() []Delimiter {
	return obj.list
}
