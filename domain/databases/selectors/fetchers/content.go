package fetchers

import "github.com/steve-care-software/webx/domain/databases/entities"

type content struct {
	recursive entities.Identifier
	selector  entities.Identifier
}

func createContentWithRecursive(
	recursive entities.Identifier,
) Content {
	return createContentInternally(recursive, nil)
}

func createContentWithSelector(
	selector entities.Identifier,
) Content {
	return createContentInternally(nil, selector)
}

func createContentInternally(
	recursive entities.Identifier,
	selector entities.Identifier,
) Content {
	out := content{
		recursive: recursive,
		selector:  selector,
	}

	return &out
}

// IsRecursive returns true if recursive, false otherwise
func (obj *content) IsRecursive() bool {
	return obj.recursive != nil
}

// Recursive returns the recursive, if any
func (obj *content) Recursive() entities.Identifier {
	return obj.recursive
}

// IsSelector returns true if selector, false otherwise
func (obj *content) IsSelector() bool {
	return obj.selector != nil
}

// Selector returns the selector, if any
func (obj *content) Selector() entities.Identifier {
	return obj.selector
}
