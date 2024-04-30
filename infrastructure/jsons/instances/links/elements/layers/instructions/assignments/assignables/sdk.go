package assignables

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables"
	json_bytes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
)

// NewAdapter creates a new adapter
func NewAdapter() assignables.Adapter {
	bytesAdapter := json_bytes.NewAdapter()
	compilerAdapter := json_compiler.NewAdapter()
	constantAdapter := json_constants.NewAdapter()
	cryptographyAdapter := json_cryptography.NewAdapter()
	builder := assignables.NewBuilder()
	return createAdapter(
		bytesAdapter.(*json_bytes.Adapter),
		compilerAdapter.(*json_compiler.Adapter),
		constantAdapter.(*json_constants.Adapter),
		cryptographyAdapter.(*json_cryptography.Adapter),
		builder,
	)
}
