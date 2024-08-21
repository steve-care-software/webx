package parameters

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"

// Parameter represents an execution parameter
type Parameter interface {
	Element() elements.Element
	Index() uint
	Name() string
}
