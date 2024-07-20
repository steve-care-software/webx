package instances

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/interruptions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/success"
	success_output "github.com/steve-care-software/datastencil/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables"
	bytes_domain "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	keys_encryptions "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	keys_encryptions_decrypts "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	keys_encryptions_encrypts "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	keys_signatures "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	keys_signatures_signs "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	keys_signatures_signs_creates "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	keys_signatures_signs_validates "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	keys_signatures_votes "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	keys_signatures_votes_creates "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	keys_signatures_votes_validates "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	assignables_executions "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	executes_inputs "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	instructions_executions "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	instructions_lists "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/layers/references"
)

// add executions

func TestAdapter_Success(t *testing.T) {
	instances := map[string]instances.Instance{
		"layers.instructions.assignments.assignables.bytes": bytes_domain.NewBytesWithHashBytesForTests(
			"myInput",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.compiles": compilers.NewCompilerWithCompileForTests(
			"myCompile",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.constants{single}": constants.NewConstantWithBoolForTests(true).(instances.Instance),
		"layers.instructions.assignments.assignables.constants{list}": constants.NewConstantsForTests([]constants.Constant{
			constants.NewConstantWithBoolForTests(true),
			constants.NewConstantWithStringForTests("this is a string"),
		}).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.decrypts": decrypts.NewDecryptForTests(
			"myCipher",
			"myPassword",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.encrypts": encrypts.NewEncryptForTests(
			"myMessage",
			"myPassword",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.encryptions.decrypts": keys_encryptions_decrypts.NewDecryptForTests(
			"myCipher",
			"myPK",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.encryptions.encrypts": keys_encryptions_encrypts.NewEncryptForTests(
			"myMessage",
			"myPubKey",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.encryptions": keys_encryptions.NewEncryptionWithGeneratePrivateKeyForTests().(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures.signs.creates": keys_signatures_signs_creates.NewCreateForTests(
			"myMessage",
			"myPK",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures.signs.validates": keys_signatures_signs_validates.NewValidateForTests(
			"mySig",
			"myMessage",
			"myPubKey",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures.signs": keys_signatures_signs.NewSignWithCreateForTests(
			keys_signatures_signs_creates.NewCreateForTests(
				"myMessage",
				"myPK",
			),
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures.votes.creates": keys_signatures_votes_creates.NewCreateForTests(
			"myMessage",
			"myRing",
			"myPK",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures.votes.validates": keys_signatures_votes_validates.NewValidateForTests(
			"myVote",
			"myMessage",
			"myHAshedRing",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures.votes": keys_signatures_votes.NewVoteWithCreateForTests(
			keys_signatures_votes_creates.NewCreateForTests(
				"myMessage",
				"myRing",
				"myPK",
			),
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys.signatures": keys_signatures.NewSignatureWithGeneratePrivateKeyForTests().(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography.keys": keys.NewKeyWithSignatureForTests(
			keys_signatures.NewSignatureWithGeneratePrivateKeyForTests(),
		).(instances.Instance),
		"layers.instructions.assignments.assignables.cryptography": cryptography.NewCryptographyWithEncryptForTests(
			encrypts.NewEncryptForTests(
				"myMessage",
				"myPassword",
			),
		).(instances.Instance),
		"layers.instructions.assignments.assignables.executions.executes.inputs": executes_inputs.NewInputWithPathForTests("myPath").(instances.Instance),
		"layers.instructions.assignments.assignables.executions.executes": executes.NewExecuteForTests(
			"myContext",
			executes_inputs.NewInputWithPathForTests("myPath"),
		).(instances.Instance),
		"layers.instructions.assignments.assignables.executions.inits":     inits.NewInitForTests("myPath", "myNme", "myDescription", "myContext").(instances.Instance),
		"layers.instructions.assignments.assignables.executions.retrieves": retrieves.NewRetrieveForTests("myContext", "myIndex", "myReturn").(instances.Instance),
		"layers.instructions.assignments.assignables.executions": assignables_executions.NewExecutionForTests(
			"myExecutable",
			assignables_executions.NewContentWithListForTests(),
		).(instances.Instance),
		"layers.instructions.assignments.assignables.lists.fetches": fetches.NewFetchForTests(
			"myList",
			"myIndex",
		).(instances.Instance),
		"layers.instructions.assignments.assignables.lists": assignables.NewAssignableWithListForTests(
			lists.NewListWithFetchForTests(
				fetches.NewFetchForTests(
					"myList",
					"myIndex",
				),
			),
		).(instances.Instance),
		"layers.instructions.assignments.assignables": lists.NewListWithFetchForTests(
			fetches.NewFetchForTests(
				"myList",
				"myIndex",
			),
		).(instances.Instance),
		"layers.instructions.assignments": assignments.NewAssignmentForTests(
			"myName",
			assignables.NewAssignableWithBytesForTests(
				bytes_domain.NewBytesWithHashBytesForTests(
					"myInput",
				),
			),
		).(instances.Instance),
		"layers.instructions.executions.merges": merges.NewMergeForTests("myBase", "myTop").(instances.Instance),
		"layers.instructions.executions": instructions_executions.NewExecutionWithMergeForTests(
			merges.NewMergeForTests("myBase", "myTop"),
		).(instances.Instance),
		"layers.instructions.lists.deletes": deletes.NewDeleteForTests(
			"myList",
			"myIndex",
		).(instances.Instance),
		"layers.instructions.lists.inserts": inserts.NewInsertForTests(
			"myList",
			"myElement",
		).(instances.Instance),
		"layers.instructions.lists": instructions_lists.NewListWithDeleteForTests(
			deletes.NewDeleteForTests(
				"myList",
				"myIndex",
			),
		).(instances.Instance),
		"layers.instructions{single}": instructions.NewInstructionWithExecutionForTests(
			instructions_executions.NewExecutionWithCommitForTests("myCommit"),
		).(instances.Instance),
		"layers.instructions{list}": instructions.NewInstructionsForTests([]instructions.Instruction{
			instructions.NewInstructionWithExecutionForTests(
				instructions_executions.NewExecutionWithCommitForTests("myCommit"),
			),
		}).(instances.Instance),
		"layers.outputs.kinds": kinds.NewKindWithContinueForTests().(instances.Instance),
		"layers.references{single}": references.NewReferenceForTests(
			"myVariable",
			[]string{"this", "is", "a", "path"},
		).(instances.Instance),
		"layers.references{list}": references.NewReferencesForTests([]references.Reference{
			references.NewReferenceForTests(
				"myVariable",
				[]string{"this", "is", "a", "path"},
			),
		}).(instances.Instance),
		"layers": layers.NewLayerWithReferencesForTests(
			instructions.NewInstructionsForTests([]instructions.Instruction{
				instructions.NewInstructionWithAssignmentForTests(
					assignments.NewAssignmentForTests(
						"anotherName",
						assignables.NewAssignableWithBytesForTests(
							bytes_domain.NewBytesWithHashBytesForTests(
								"anotherInput",
							),
						),
					),
				),
				instructions.NewInstructionWithRaiseErrorForTests(22),
				instructions.NewInstructionWithStopForTests(),
			}),
			outputs.NewOutputForTests(
				"myVariable",
				kinds.NewKindWithContinueForTests(),
			),
			"myInput",
			references.NewReferencesForTests([]references.Reference{
				references.NewReferenceForTests(
					"myVariable",
					[]string{"this", "is", "a", "path"},
				),
			}),
		).(instances.Instance),
		"layers.outputs": outputs.NewOutputForTests(
			"myVariable",
			kinds.NewKindWithContinueForTests(),
		).(instances.Instance),
		"executions.results.interruptions.failures": failures.NewFailureForTests(
			uint(34),
			uint(32),
			false,
		).(instances.Instance),
		"executions.results.interruptions": interruptions.NewInterruptionWithStopForTests(
			23,
		).(instances.Instance),
		"executions.results.success.output": success_output.NewOutputForTests(
			[]byte("this is an input"),
		).(instances.Instance),
		"executions.results.success": success.NewSuccessForTests(
			success_output.NewOutputForTests(
				[]byte("this is an input"),
			),
			kinds.NewKindWithPromptForTests(),
		).(instances.Instance),
		"executions.results": results.NewResultWithSuccessForTests(
			success.NewSuccessForTests(
				success_output.NewOutputForTests(
					[]byte("this is an input"),
				),
				kinds.NewKindWithPromptForTests(),
			),
		),
		"executions{single}": executions.NewExecutionForTests(
			[]byte("myInput"),
			layers.NewLayerWithReferencesForTests(
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithAssignmentForTests(
						assignments.NewAssignmentForTests(
							"anotherName",
							assignables.NewAssignableWithBytesForTests(
								bytes_domain.NewBytesWithHashBytesForTests(
									"anotherInput",
								),
							),
						),
					),
					instructions.NewInstructionWithRaiseErrorForTests(22),
					instructions.NewInstructionWithStopForTests(),
				}),
				outputs.NewOutputForTests(
					"myVariable",
					kinds.NewKindWithContinueForTests(),
				),
				"myInput",
				references.NewReferencesForTests([]references.Reference{
					references.NewReferenceForTests(
						"myVariable",
						[]string{"this", "is", "a", "path"},
					),
				}),
			),
			results.NewResultWithSuccessForTests(
				success.NewSuccessForTests(
					success_output.NewOutputForTests(
						[]byte("this is an input"),
					),
					kinds.NewKindWithPromptForTests(),
				),
			),
		),
		"executions{list}": executions.NewExecutionsForTests([]executions.Execution{
			executions.NewExecutionForTests(
				[]byte("myInput"),
				layers.NewLayerWithReferencesForTests(
					instructions.NewInstructionsForTests([]instructions.Instruction{
						instructions.NewInstructionWithAssignmentForTests(
							assignments.NewAssignmentForTests(
								"anotherName",
								assignables.NewAssignableWithBytesForTests(
									bytes_domain.NewBytesWithHashBytesForTests(
										"anotherInput",
									),
								),
							),
						),
						instructions.NewInstructionWithRaiseErrorForTests(22),
						instructions.NewInstructionWithStopForTests(),
					}),
					outputs.NewOutputForTests(
						"myVariable",
						kinds.NewKindWithContinueForTests(),
					),
					"myInput",
					references.NewReferencesForTests([]references.Reference{
						references.NewReferenceForTests(
							"myVariable",
							[]string{"this", "is", "a", "path"},
						),
					}),
				),
				results.NewResultWithSuccessForTests(
					success.NewSuccessForTests(
						success_output.NewOutputForTests(
							[]byte("this is an input"),
						),
						kinds.NewKindWithPromptForTests(),
					),
				),
			),
		}),
	}

	adapter := NewAdapter()

	for keyname, oneInstance := range instances {
		retBytes, err := adapter.ToBytes(oneInstance)
		if err != nil {
			t.Errorf("keyname: %s, the error was expected to be nil, error returned: %s", keyname, err.Error())
			return
		}

		retIns, err := adapter.ToInstance(retBytes)
		if err != nil {
			t.Errorf("keyname: %s, the error was expected to be nil, error returned: %s", keyname, err.Error())
			return
		}

		if !bytes.Equal(oneInstance.Hash().Bytes(), retIns.Hash().Bytes()) {
			t.Errorf("keyname: %s, the returned instance is invalid", keyname)
			return
		}

		if !reflect.DeepEqual(oneInstance, retIns) {
			t.Errorf("keyname; %s, the returned instance is invalid", keyname)
		}
	}

}
