package instructions

import (
	"testing"

	application_assignments "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments"
	application_assignments_assignables "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables"
	application_assignments_assignables_bytes "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/bytes"
	application_assignments_compilers "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/compilers"
	application_assignments_assignables_constants "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/constants"
	application_assignments_assignables_cryptography "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/cryptography"
	application_assignments_assignables_cryptography_decrypts "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/cryptography/decrypts"
	application_assignments_assignables_cryptography_encrypts "github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
)

func TestExecute_WithIsStop_Success(t *testing.T) {
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithStopForTests(),
	})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	stack := stacks.NewFactory().Create()
	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsStop() {
		t.Errorf("the instructions were expected to be stopped")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithRaisedError_Success(t *testing.T) {
	raisedError := uint(45)
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithRaiseErrorForTests(raisedError),
	})

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	stack := stacks.NewFactory().Create()
	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be a failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithAssignment_Succeeds_Success(t *testing.T) {
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						firstVariable,
						stacks.NewAssignableWithBytesForTests(
							firstValue,
						),
					),
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	retStackAssignable, err := retStack.Head().Fetch(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !retStackAssignable.IsBytes() {
		t.Errorf("the returned assignable was expected to contain bytes")
		return
	}
}

func TestExecute_WithAssignment_Fails_ReturnsFailure(t *testing.T) {
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	firstVariable := "firstVar"
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be a failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	failure := retInterruption.Failure()
	if !failure.HasMessage() {
		t.Errorf("the failure was expected to contain a message")
		return
	}
}

func TestExecute_WithCondition_ConditionIsFalse_WithAssignment_Succeeds_Success(t *testing.T) {
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := false

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						firstVariable,
						stacks.NewAssignableWithBytesForTests(
							firstValue,
						),
					),
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}

	retStackAssignable, err := retStack.Head().Fetch(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !retStackAssignable.IsBytes() {
		t.Errorf("the returned assignable was expected to contain bytes")
		return
	}
}

func TestExecute_WithCondition_ConditionIsTrue_WithAssignment_ExecutesConditionInstructions_Stops_ReturnsInterruption(t *testing.T) {
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := true

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						firstVariable,
						stacks.NewAssignableWithBytesForTests(
							firstValue,
						),
					),
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithStopForTests(),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsStop() {
		t.Errorf("the instructions were expected to be stopped")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithCondition_ConditionIsTrue_WithAssignment_ExecutesConditionInstructions_Failure_ReturnsInterruption(t *testing.T) {
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := true

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						firstVariable,
						stacks.NewAssignableWithBytesForTests(
							firstValue,
						),
					),
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithConditionForTests(
			instructions.NewConditionForTest(
				conditionVar,
				instructions.NewInstructionsForTests([]instructions.Instruction{
					instructions.NewInstructionWithRaiseErrorForTests(45),
				}),
			),
		),
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption == nil {
		t.Errorf("the interruption was expected to be valid")
		return
	}

	if !retInterruption.IsFailure() {
		t.Errorf("the instructions were expected to be failure")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}

func TestExecute_WithCondition_ConditionIsTrue_WithAssignment_ExecutesConditionInstructions_Succeeds_Success(t *testing.T) {
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		application_assignments.NewApplication(
			application_assignments_assignables.NewApplication(
				application_assignments_compilers.NewApplication(
					instanceAdapter,
				),
				application_assignments_assignables_bytes.NewApplication(),
				application_assignments_assignables_constants.NewApplication(),
				application_assignments_assignables_cryptography.NewApplication(
					application_assignments_assignables_cryptography_decrypts.NewApplication(
						encryptor,
					),
					application_assignments_assignables_cryptography_encrypts.NewApplication(
						encryptor,
					), nil,
				),
			),
		),
	)

	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	conditionVar := "myCondition"
	conditionValue := true

	stack := stacks.NewStackForTests(
		stacks.NewFramesForTests([]stacks.Frame{
			stacks.NewFrameWithAssignmentsForTests(
				stacks.NewAssignmentsForTests([]stacks.Assignment{
					stacks.NewAssignmentForTests(
						firstVariable,
						stacks.NewAssignableWithBytesForTests(
							firstValue,
						),
					),
					stacks.NewAssignmentForTests(
						secondVariable,
						stacks.NewAssignableWithBytesForTests(
							secondValue,
						),
					),
					stacks.NewAssignmentForTests(
						conditionVar,
						stacks.NewAssignableWithBoolForTests(
							conditionValue,
						),
					),
				}),
			),
		}),
	)

	name := "myName"
	instructions := instructions.NewInstructionsForTests([]instructions.Instruction{
		instructions.NewInstructionWithAssignmentForTests(
			assignments.NewAssignmentForTests(
				name,
				assignables.NewAssignableWithBytesForTests(
					assignable_bytes.NewBytesWithJoinForTests([]string{
						firstVariable,
						secondVariable,
					}),
				),
			),
		),
	})

	retStack, retInterruption, err := application.Execute(stack, instructions)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retInterruption != nil {
		t.Errorf("the interruption was expected to NOT be valid")
		return
	}

	if retStack == nil {
		t.Errorf("the stack was expected to be valid")
		return
	}
}
