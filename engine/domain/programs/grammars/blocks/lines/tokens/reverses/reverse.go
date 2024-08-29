package reverses

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"

type reverse struct {
	escape elements.Element
}

func createReverse() Reverse {
	return createReverseInternally(nil)
}

func createReverseWithEscape(
	escape elements.Element,
) Reverse {
	return createReverseInternally(escape)
}

func createReverseInternally(
	escape elements.Element,
) Reverse {
	out := reverse{
		escape: escape,
	}

	return &out
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *reverse) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *reverse) Escape() elements.Element {
	return obj.escape
}
