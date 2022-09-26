package methods

import "github.com/steve-care-software/syntax/domain/programs/instructions/applications/modules"

// Application represents the module application
type Application interface {
	Combines() modules.Module
	Compiles() modules.Module
	Executes() modules.Module
	Extracts() modules.Module
	Programs() modules.Module
	Tokenizes() modules.Module
}
