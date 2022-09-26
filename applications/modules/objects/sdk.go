package objects

import "github.com/steve-care-software/syntax/domain/programs/instructions/applications/modules"

// Application represents the object application
type Application interface {
	Compilers() modules.Module
	Criterias() modules.Module
	Instructions() modules.Module
	Grammars() modules.Module
}
