package instances

import (
	"errors"

	"github.com/steve-care-software/webx/engine/stencils/domain/instances"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions/failures"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/success"
	success_output "github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	keys_encryptions "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	keys_encryptions_decrypts "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	keys_encryptions_encrypts "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	keys_signatures "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	keys_signatures_signs "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	keys_signatures_signs_creates "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	keys_signatures_signs_validates "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	keys_signatures_votes "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	keys_signatures_votes_creates "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	keys_signatures_votes_validates "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	assignables_executions "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	executes_inputs "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	instructions_executions "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/executions/merges"
	instructions_lists "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/lists"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/lists/deletes"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/lists/inserts"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs/kinds"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/references"
)

type adapter struct {
	failureAdapter                   failures.Adapter
	interruptionAdapter              interruptions.Adapter
	successOutputAdapter             success_output.Adapter
	successAdapter                   success.Adapter
	resultAdapter                    results.Adapter
	bytesAdapter                     bytes_domain.Adapter
	compilerAdapter                  compilers.Adapter
	constantAdapter                  constants.Adapter
	decryptAdapter                   decrypts.Adapter
	encryptAdapter                   encrypts.Adapter
	keysEncDecryptAdapter            keys_encryptions_decrypts.Adapter
	keysEncEncryptAdapter            keys_encryptions_encrypts.Adapter
	keysEncAdapter                   keys_encryptions.Adapter
	keysSignatureSignCreateAdapter   keys_signatures_signs_creates.Adapter
	keysSignatureSignValidateAdapter keys_signatures_signs_validates.Adapter
	keysSignatureSignAdapter         keys_signatures_signs.Adapter
	keysSignatureVoteCreateAdapter   keys_signatures_votes_creates.Adapter
	keysSignatureVoteValidateAdapter keys_signatures_votes_validates.Adapter
	keysSignatureVoteAdapter         keys_signatures_votes.Adapter
	keysSignatureAdapter             keys_signatures.Adapter
	keyAdapter                       keys.Adapter
	cryptographyAdapter              cryptography.Adapter
	executeInputAdapter              executes_inputs.Adapter
	executeAdapter                   executes.Adapter
	initAdapter                      inits.Adapter
	retrieveAdapter                  retrieves.Adapter
	assignableExecutionAdapter       assignables_executions.Adapter
	fetchesAdapter                   fetches.Adapter
	listAdapter                      lists.Adapter
	assignableAdapter                assignables.Adapter
	assignmentAdapter                assignments.Adapter
	mergeAdapter                     merges.Adapter
	instructionExecutionAdapter      instructions_executions.Adapter
	deleteAdapter                    deletes.Adapter
	insertAdapter                    inserts.Adapter
	instructionListAdapter           instructions_lists.Adapter
	instructionAdapter               instructions.Adapter
	kindAdapter                      kinds.Adapter
	outputAdapter                    outputs.Adapter
	referenceAdapter                 references.Adapter
	layerAdapter                     layers.Adapter
	executionAdapter                 executions.Adapter
}

func createAdapter(
	failureAdapter failures.Adapter,
	interruptionAdapter interruptions.Adapter,
	successOutputAdapter success_output.Adapter,
	successAdapter success.Adapter,
	resultAdapter results.Adapter,
	bytesAdapter bytes_domain.Adapter,
	compilerAdapter compilers.Adapter,
	constantAdapter constants.Adapter,
	decryptAdapter decrypts.Adapter,
	encryptAdapter encrypts.Adapter,
	keysEncDecryptAdapter keys_encryptions_decrypts.Adapter,
	keysEncEncryptAdapter keys_encryptions_encrypts.Adapter,
	keysEncAdapter keys_encryptions.Adapter,
	keysSignatureSignCreateAdapter keys_signatures_signs_creates.Adapter,
	keysSignatureSignValidateAdapter keys_signatures_signs_validates.Adapter,
	keysSignatureSignAdapter keys_signatures_signs.Adapter,
	keysSignatureVoteCreateAdapter keys_signatures_votes_creates.Adapter,
	keysSignatureVoteValidateAdapter keys_signatures_votes_validates.Adapter,
	keysSignatureVoteAdapter keys_signatures_votes.Adapter,
	keysSignatureAdapter keys_signatures.Adapter,
	keyAdapter keys.Adapter,
	cryptographyAdapter cryptography.Adapter,
	executeInputAdapter executes_inputs.Adapter,
	executeAdapter executes.Adapter,
	initAdapter inits.Adapter,
	retrieveAdapter retrieves.Adapter,
	assignableExecutionAdapter assignables_executions.Adapter,
	fetchesAdapter fetches.Adapter,
	listAdapter lists.Adapter,
	assignableAdapter assignables.Adapter,
	assignmentAdapter assignments.Adapter,
	mergeAdapter merges.Adapter,
	instructionExecutionAdapter instructions_executions.Adapter,
	deleteAdapter deletes.Adapter,
	insertAdapter inserts.Adapter,
	instructionListAdapter instructions_lists.Adapter,
	instructionAdapter instructions.Adapter,
	kindAdapter kinds.Adapter,
	outputAdapter outputs.Adapter,
	referenceAdapter references.Adapter,
	layerAdapter layers.Adapter,
	executionAdapter executions.Adapter,
) instances.Adapter {
	out := adapter{
		failureAdapter:                   failureAdapter,
		interruptionAdapter:              interruptionAdapter,
		successOutputAdapter:             successOutputAdapter,
		successAdapter:                   successAdapter,
		resultAdapter:                    resultAdapter,
		bytesAdapter:                     bytesAdapter,
		compilerAdapter:                  compilerAdapter,
		constantAdapter:                  constantAdapter,
		decryptAdapter:                   decryptAdapter,
		encryptAdapter:                   encryptAdapter,
		keysEncDecryptAdapter:            keysEncDecryptAdapter,
		keysEncEncryptAdapter:            keysEncEncryptAdapter,
		keysEncAdapter:                   keysEncAdapter,
		keysSignatureSignCreateAdapter:   keysSignatureSignCreateAdapter,
		keysSignatureSignValidateAdapter: keysSignatureSignValidateAdapter,
		keysSignatureSignAdapter:         keysSignatureSignAdapter,
		keysSignatureVoteCreateAdapter:   keysSignatureVoteCreateAdapter,
		keysSignatureVoteValidateAdapter: keysSignatureVoteValidateAdapter,
		keysSignatureVoteAdapter:         keysSignatureVoteAdapter,
		keysSignatureAdapter:             keysSignatureAdapter,
		keyAdapter:                       keyAdapter,
		cryptographyAdapter:              cryptographyAdapter,
		executeInputAdapter:              executeInputAdapter,
		executeAdapter:                   executeAdapter,
		initAdapter:                      initAdapter,
		retrieveAdapter:                  retrieveAdapter,
		assignableExecutionAdapter:       assignableExecutionAdapter,
		fetchesAdapter:                   fetchesAdapter,
		listAdapter:                      listAdapter,
		assignableAdapter:                assignableAdapter,
		assignmentAdapter:                assignmentAdapter,
		mergeAdapter:                     mergeAdapter,
		instructionExecutionAdapter:      instructionExecutionAdapter,
		deleteAdapter:                    deleteAdapter,
		insertAdapter:                    insertAdapter,
		instructionListAdapter:           instructionListAdapter,
		instructionAdapter:               instructionAdapter,
		kindAdapter:                      kindAdapter,
		outputAdapter:                    outputAdapter,
		referenceAdapter:                 referenceAdapter,
		layerAdapter:                     layerAdapter,
		executionAdapter:                 executionAdapter,
	}

	return &out
}

// ToBytes converts an instance to bytes
func (app *adapter) ToBytes(ins instances.Instance) ([]byte, error) {
	if casted, ok := ins.(kinds.Kind); ok {
		return app.kindAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(outputs.Output); ok {
		return app.outputAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(failures.Failure); ok {
		return app.failureAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(interruptions.Interruption); ok {
		return app.interruptionAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(success_output.Output); ok {
		return app.successOutputAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(success.Success); ok {
		return app.successAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(results.Result); ok {
		return app.resultAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(bytes_domain.Bytes); ok {
		return app.bytesAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(compilers.Compiler); ok {
		return app.compilerAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(constants.Constant); ok {
		return app.constantAdapter.InstanceToBytes(casted)
	}

	if casted, ok := ins.(constants.Constants); ok {
		return app.constantAdapter.InstancesToBytes(casted)
	}

	if casted, ok := ins.(decrypts.Decrypt); ok {
		return app.decryptAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(encrypts.Encrypt); ok {
		return app.encryptAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures_votes_creates.Create); ok {
		return app.keysSignatureVoteCreateAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures_votes_validates.Validate); ok {
		return app.keysSignatureVoteValidateAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures_signs_creates.Create); ok {
		return app.keysSignatureSignCreateAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures_signs_validates.Validate); ok {
		return app.keysSignatureSignValidateAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_encryptions_decrypts.Decrypt); ok {
		return app.keysEncDecryptAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_encryptions_encrypts.Encrypt); ok {
		return app.keysEncEncryptAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_encryptions.Encryption); ok {
		return app.keysEncAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures_signs.Sign); ok {
		return app.keysSignatureSignAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures_votes.Vote); ok {
		return app.keysSignatureVoteAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys_signatures.Signature); ok {
		return app.keysSignatureAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(keys.Key); ok {
		return app.keyAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(cryptography.Cryptography); ok {
		return app.cryptographyAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(retrieves.Retrieve); ok {
		return app.retrieveAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(executes.Execute); ok {
		return app.executeAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(inits.Init); ok {
		return app.initAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(executes_inputs.Input); ok {
		return app.executeInputAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(assignables_executions.Execution); ok {
		return app.assignableExecutionAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(fetches.Fetch); ok {
		return app.fetchesAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(deletes.Delete); ok {
		return app.deleteAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(inserts.Insert); ok {
		return app.insertAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(lists.List); ok {
		return app.listAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(assignables.Assignable); ok {
		return app.assignableAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(assignments.Assignment); ok {
		return app.assignmentAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(merges.Merge); ok {
		return app.mergeAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(instructions_executions.Execution); ok {
		return app.instructionExecutionAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(instructions_lists.List); ok {
		return app.instructionListAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(instructions.Instruction); ok {
		return app.instructionAdapter.InstanceToBytes(casted)
	}

	if casted, ok := ins.(instructions.Instructions); ok {
		return app.instructionAdapter.InstancesToBytes(casted)
	}

	if casted, ok := ins.(references.Reference); ok {
		return app.referenceAdapter.InstanceToBytes(casted)
	}

	if casted, ok := ins.(references.References); ok {
		return app.referenceAdapter.InstancesToBytes(casted)
	}

	if casted, ok := ins.(layers.Layer); ok {
		return app.layerAdapter.ToBytes(casted)
	}

	if casted, ok := ins.(executions.Execution); ok {
		return app.executionAdapter.InstanceToBytes(casted)
	}

	if casted, ok := ins.(executions.Executions); ok {
		return app.executionAdapter.InstancesToBytes(casted)
	}

	return nil, errors.New("the Instance could not be converted to bytes")
}

// ToInstance converts bytes to instance
func (app *adapter) ToInstance(data []byte) (instances.Instance, error) {
	kind, err := app.kindAdapter.ToInstance(data)
	if err == nil {
		return kind, nil
	}

	output, err := app.outputAdapter.ToInstance(data)
	if err == nil {
		return output, nil
	}

	failure, err := app.failureAdapter.ToInstance(data)
	if err == nil {
		return failure, nil
	}

	interruption, err := app.interruptionAdapter.ToInstance(data)
	if err == nil {
		return interruption, nil
	}

	successOutput, err := app.successOutputAdapter.ToInstance(data)
	if err == nil {
		return successOutput, nil
	}

	success, err := app.successAdapter.ToInstance(data)
	if err == nil {
		return success, nil
	}

	result, err := app.resultAdapter.ToInstance(data)
	if err == nil {
		return result, nil
	}

	retBytes, err := app.bytesAdapter.ToInstance(data)
	if err == nil {
		return retBytes, nil
	}

	compiler, err := app.compilerAdapter.ToInstance(data)
	if err == nil {
		return compiler, nil
	}

	constant, err := app.constantAdapter.BytesToInstance(data)
	if err == nil {
		return constant, nil
	}

	constants, err := app.constantAdapter.BytesToInstances(data)
	if err == nil {
		return constants, nil
	}

	decrypt, err := app.decryptAdapter.ToInstance(data)
	if err == nil {
		return decrypt, nil
	}

	encrypt, err := app.encryptAdapter.ToInstance(data)
	if err == nil {
		return encrypt, nil
	}

	keysSigVoteCreate, err := app.keysSignatureVoteCreateAdapter.ToInstance(data)
	if err == nil {
		return keysSigVoteCreate, nil
	}

	keySigVoteValidate, err := app.keysSignatureVoteValidateAdapter.ToInstance(data)
	if err == nil {
		return keySigVoteValidate, nil
	}

	keysSigSignCreate, err := app.keysSignatureSignCreateAdapter.ToInstance(data)
	if err == nil {
		return keysSigSignCreate, nil
	}

	keysSigValidate, err := app.keysSignatureSignValidateAdapter.ToInstance(data)
	if err == nil {
		return keysSigValidate, nil
	}

	keysEncDecrypt, err := app.keysEncDecryptAdapter.ToInstance(data)
	if err == nil {
		return keysEncDecrypt, nil
	}

	keysEncEncrypt, err := app.keysEncEncryptAdapter.ToInstance(data)
	if err == nil {
		return keysEncEncrypt, nil
	}

	keysEnc, err := app.keysEncAdapter.ToInstance(data)
	if err == nil {
		return keysEnc, nil
	}

	keysSigSign, err := app.keysSignatureSignAdapter.ToInstance(data)
	if err == nil {
		return keysSigSign, nil
	}

	keysSigVote, err := app.keysSignatureVoteAdapter.ToInstance(data)
	if err == nil {
		return keysSigVote, nil
	}

	keysSig, err := app.keysSignatureAdapter.ToInstance(data)
	if err == nil {
		return keysSig, nil
	}

	key, err := app.keyAdapter.ToInstance(data)
	if err == nil {
		return key, nil
	}

	cryptography, err := app.cryptographyAdapter.ToInstance(data)
	if err == nil {
		return cryptography, nil
	}

	retrieve, err := app.retrieveAdapter.ToInstance(data)
	if err == nil {
		return retrieve, nil
	}

	execute, err := app.executeAdapter.ToInstance(data)
	if err == nil {
		return execute, nil
	}

	init, err := app.initAdapter.ToInstance(data)
	if err == nil {
		return init, nil
	}

	executeInput, err := app.executeInputAdapter.ToInstance(data)
	if err == nil {
		return executeInput, nil
	}

	assignableExecution, err := app.assignableExecutionAdapter.ToInstance(data)
	if err == nil {
		return assignableExecution, nil
	}

	fetch, err := app.fetchesAdapter.ToInstance(data)
	if err == nil {
		return fetch, nil
	}

	del, err := app.deleteAdapter.ToInstance(data)
	if err == nil {
		return del, nil
	}

	insert, err := app.insertAdapter.ToInstance(data)
	if err == nil {
		return insert, nil
	}

	list, err := app.listAdapter.ToInstance(data)
	if err == nil {
		return list, nil
	}

	assignable, err := app.assignableAdapter.ToInstance(data)
	if err == nil {
		return assignable, nil
	}

	assignment, err := app.assignmentAdapter.ToInstance(data)
	if err == nil {
		return assignment, nil
	}

	merge, err := app.mergeAdapter.ToInstance(data)
	if err == nil {
		return merge, nil
	}

	instructionExecution, err := app.instructionExecutionAdapter.ToInstance(data)
	if err == nil {
		return instructionExecution, nil
	}

	instructionList, err := app.instructionListAdapter.ToInstance(data)
	if err == nil {
		return instructionList, nil
	}

	instruction, err := app.instructionAdapter.BytesToInstance(data)
	if err == nil {
		return instruction, nil
	}

	instructions, err := app.instructionAdapter.BytesToInstances(data)
	if err == nil {
		return instructions, nil
	}

	reference, err := app.referenceAdapter.BytesToInstance(data)
	if err == nil {
		return reference, nil
	}

	references, err := app.referenceAdapter.BytesToInstances(data)
	if err == nil {
		return references, nil
	}

	layer, err := app.layerAdapter.ToInstance(data)
	if err == nil {
		return layer, nil
	}

	execution, err := app.executionAdapter.BytesToInstance(data)
	if err == nil {
		return execution, nil
	}

	executions, err := app.executionAdapter.BytesToInstances(data)
	if err == nil {
		return executions, nil
	}

	return nil, errors.New("the bytes could not be converted to an Instance")
}
