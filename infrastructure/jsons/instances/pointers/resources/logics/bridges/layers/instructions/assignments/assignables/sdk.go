package assignables

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables"
	json_bytes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography"
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases"
	json_lists "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/lists"
)

// NewAdapter creates a new adapter
func NewAdapter() assignables.Adapter {
	bytesAdapter := json_bytes.NewAdapter()
	compilerAdapter := json_compiler.NewAdapter()
	constantAdapter := json_constants.NewAdapter()
	cryptographyAdapter := json_cryptography.NewAdapter()
	databaseAdapter := json_databases.NewAdapter()
	listAdapter := json_lists.NewAdapter()
	builder := assignables.NewBuilder()
	return createAdapter(
		bytesAdapter.(*json_bytes.Adapter),
		compilerAdapter.(*json_compiler.Adapter),
		constantAdapter.(*json_constants.Adapter),
		cryptographyAdapter.(*json_cryptography.Adapter),
		databaseAdapter.(*json_databases.Adapter),
		listAdapter.(*json_lists.Adapter),
		builder,
	)
}
