package assignables

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"
	json_bytes "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography"
	json_executables "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executables"
	json_executions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions"
	json_lists "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists"
)

// NewAdapter creates a new adapter
func NewAdapter() assignables.Adapter {
	executionAdapter := json_executions.NewAdapter()
	bytesAdapter := json_bytes.NewAdapter()
	compilerAdapter := json_compiler.NewAdapter()
	constantAdapter := json_constants.NewAdapter()
	cryptographyAdapter := json_cryptography.NewAdapter()
	listAdapter := json_lists.NewAdapter()
	executableAdapter := json_executables.NewAdapter()
	builder := assignables.NewBuilder()
	return createAdapter(
		executionAdapter.(*json_executions.Adapter),
		bytesAdapter.(*json_bytes.Adapter),
		compilerAdapter.(*json_compiler.Adapter),
		constantAdapter.(*json_constants.Adapter),
		cryptographyAdapter.(*json_cryptography.Adapter),
		listAdapter.(*json_lists.Adapter),
		executableAdapter.(*json_executables.Adapter),
		builder,
	)
}
