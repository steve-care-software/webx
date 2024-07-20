package applications

import "github.com/steve-care-software/webx/engine/states/applications"

// NewBuilder creates a new application builder
func NewBuilder() applications.Builder {
	return createBuilder()
}
