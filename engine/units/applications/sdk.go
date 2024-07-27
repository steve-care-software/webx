package applications

import (
	"github.com/steve-care-software/webx/engine/units/applications/blockchains"
	"github.com/steve-care-software/webx/engine/units/applications/identities"
)

// Application represents the units database
type Application interface {
	Identity() identities.Application
	Blockchain() blockchains.Application
}
