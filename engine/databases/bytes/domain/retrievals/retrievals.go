package retrievals

type retrievals struct {
	list []Retrieval
}

func createRetrievals(
	list []Retrieval,
) Retrievals {
	out := retrievals{
		list: list,
	}

	return &out
}

// List returns the list of retrievals
func (obj *retrievals) List() []Retrieval {
	return obj.list
}
