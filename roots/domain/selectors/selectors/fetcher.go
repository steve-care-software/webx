package selectors

type fetcher struct {
	recursive string
	selector  Selector
}

func createFetcherWithRecursive(
	recursive string,
) Fetcher {
	return createFetcherInternally(recursive, nil)
}

func createFetcherWithSelector(
	selector Selector,
) Fetcher {
	return createFetcherInternally("", selector)
}

func createFetcherInternally(
	recursive string,
	selector Selector,
) Fetcher {
	out := fetcher{
		recursive: recursive,
		selector:  selector,
	}

	return &out
}

// IsRecursive returns true if recursive, false otherwise
func (obj *fetcher) IsRecursive() bool {
	return obj.recursive != ""
}

// Recursive returns the recursive selector's token name
func (obj *fetcher) Recursive() string {
	return obj.recursive
}

// IsSelector returns true if selector, false otherwise
func (obj *fetcher) IsSelector() bool {
	return obj.selector != nil
}

// Selector returns the selector if any
func (obj *fetcher) Selector() Selector {
	return obj.selector
}
