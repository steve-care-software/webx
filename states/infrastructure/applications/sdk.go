package applications

import "github.com/steve-care-software/datastencil/states/applications"

// NewBuilder creates a new application builder
func NewBuilder() applications.Builder {
	return createBuilder()
}
