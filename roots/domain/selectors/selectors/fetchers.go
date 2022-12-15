package selectors

type fetchers struct {
	list []Fetcher
}

func createFetchers(
	list []Fetcher,
) Fetchers {
	out := fetchers{
		list: list,
	}

	return &out
}

// List returns the fetchers
func (obj *fetchers) List() []Fetcher {
	return obj.list
}
