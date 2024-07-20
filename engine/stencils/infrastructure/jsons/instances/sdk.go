package instances

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances"
	json_executions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions"
	json_results "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results"
	json_interruptions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results/interruptions"
	json_failures "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results/interruptions/failures"
	json_success "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results/success"
	json_success_outputs "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/executions/results/success/outputs"
	json_layers "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers"
	json_instructions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions"
	json_assignments "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments"
	json_assignables "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables"
	json_bytes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography"
	json_decrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	json_encrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	json_keys "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys"
	json_keys_encryptions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	json_keys_encryptions_decrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	json_keys_encryptions_encrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	json_keys_signatures "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	json_keys_signatures_signs "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	json_keys_signatures_signs_creates "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	json_keys_signatures_signs_validates "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	json_keys_signatures_votes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	json_keys_signatures_votes_creates "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	json_keys_signatures_votes_validates "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	json_assignables_executions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions"
	json_executes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes"
	json_executes_inputs "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
	json_inits "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/inits"
	json_retrieves "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions/retrieves"
	json_lists "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists"
	json_fetches "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists/fetches"
	json_instructions_executions "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/executions"
	json_merges "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/executions/merges"
	json_instructions_lists "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/lists"
	json_instructions_lists_deletes "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/lists/deletes"
	json_instructions_lists_inserts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/lists/inserts"
	json_outputs "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/outputs"
	json_outputs_kinds "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/outputs/kinds"
	json_references "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/references"
)

// NewAdapter creates a new adapter
func NewAdapter() instances.Adapter {
	failureAdapter := json_failures.NewAdapter()
	interruptionAdapter := json_interruptions.NewAdapter()
	successOutputAdapter := json_success_outputs.NewAdapter()
	successAdapter := json_success.NewAdapter()
	resultAdapter := json_results.NewAdapter()
	bytesAdapter := json_bytes.NewAdapter()
	compilerAdapter := json_compiler.NewAdapter()
	constantAdapter := json_constants.NewAdapter()
	decryptAdapter := json_decrypts.NewAdapter()
	encryptAdapter := json_encrypts.NewAdapter()
	keysEncryptionDecryptAdapter := json_keys_encryptions_decrypts.NewAdapter()
	keysEncryptionEncryptAdapter := json_keys_encryptions_encrypts.NewAdapter()
	keysEncryptionAdapter := json_keys_encryptions.NewAdapter()
	keysSignatureSignCreateAdapter := json_keys_signatures_signs_creates.NewAdapter()
	keysSignatureSignValidateAdapter := json_keys_signatures_signs_validates.NewAdapter()
	keysSignatureSignAdapter := json_keys_signatures_signs.NewAdapter()
	keysSignatureVoteCreateAdapter := json_keys_signatures_votes_creates.NewAdapter()
	keysSignatureVoteValidateAdapter := json_keys_signatures_votes_validates.NewAdapter()
	keysSignatureVoteAdapter := json_keys_signatures_votes.NewAdapter()
	keysSignatureAdapter := json_keys_signatures.NewAdapter()
	keyAdapter := json_keys.NewAdapter()
	cryptographyAdapter := json_cryptography.NewAdapter()
	executeInputAdapter := json_executes_inputs.NewAdapter()
	executeAdapter := json_executes.NewAdapter()
	initAdapter := json_inits.NewAdapter()
	retrieveAdapter := json_retrieves.NewAdapter()
	assignableExecutionAdapter := json_assignables_executions.NewAdapter()
	fetchAdapter := json_fetches.NewAdapter()
	listAdapter := json_lists.NewAdapter()
	assignableAdapter := json_assignables.NewAdapter()
	assignmentAdapter := json_assignments.NewAdapter()
	mergeAdapter := json_merges.NewAdapter()
	instructionExecutionAdapter := json_instructions_executions.NewAdapter()
	listDeleteAdapter := json_instructions_lists_deletes.NewAdapter()
	listInsertAdapter := json_instructions_lists_inserts.NewAdapter()
	instructionListAdapter := json_instructions_lists.NewAdapter()
	instructionAdapter := json_instructions.NewAdapter()
	kindAdapter := json_outputs_kinds.NewAdapter()
	outputAdapter := json_outputs.NewAdapter()
	referenceAdapter := json_references.NewAdapter()
	layerAdapter := json_layers.NewAdapter()
	executionAdapter := json_executions.NewAdapter()
	return createAdapter(
		failureAdapter,
		interruptionAdapter,
		successOutputAdapter,
		successAdapter,
		resultAdapter,
		bytesAdapter,
		compilerAdapter,
		constantAdapter,
		decryptAdapter,
		encryptAdapter,
		keysEncryptionDecryptAdapter,
		keysEncryptionEncryptAdapter,
		keysEncryptionAdapter,
		keysSignatureSignCreateAdapter,
		keysSignatureSignValidateAdapter,
		keysSignatureSignAdapter,
		keysSignatureVoteCreateAdapter,
		keysSignatureVoteValidateAdapter,
		keysSignatureVoteAdapter,
		keysSignatureAdapter,
		keyAdapter,
		cryptographyAdapter,
		executeInputAdapter,
		executeAdapter,
		initAdapter,
		retrieveAdapter,
		assignableExecutionAdapter,
		fetchAdapter,
		listAdapter,
		assignableAdapter,
		assignmentAdapter,
		mergeAdapter,
		instructionExecutionAdapter,
		listDeleteAdapter,
		listInsertAdapter,
		instructionListAdapter,
		instructionAdapter,
		kindAdapter,
		outputAdapter,
		referenceAdapter,
		layerAdapter,
		executionAdapter,
	)
}
