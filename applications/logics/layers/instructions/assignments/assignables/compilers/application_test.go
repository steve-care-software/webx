package compilers

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
	"github.com/steve-care-software/datastencil/domain/stacks/mocks"
)

func TestExecute_withCompile_compileExistsInFrame_adapterSucceeds_Success(t *testing.T) {
	compileVar := "compileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				compileVar,
				stacks.NewAssignableWithBytesForTests(
					compile,
				),
			),
		}),
	)

	instruction := compilers.NewCompilerWithCompileForTests(compileVar)
	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{
			string(compile): compiledInstance,
		},
	)

	application := NewApplication(
		instanceAdapter,
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsInstance() {
		t.Errorf("the assignable was expected to contain an instance")
		return
	}

	retInstance := retAssignable.Instance()
	if !reflect.DeepEqual(compiledInstance.(instances.Instance), retInstance) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestExecute_withCompile_compileExistsInFrame_adapterFails_returnsError(t *testing.T) {
	compileVar := "compileVar"
	compile := []byte("this is some data")
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				compileVar,
				stacks.NewAssignableWithBytesForTests(
					compile,
				),
			),
		}),
	)

	instruction := compilers.NewCompilerWithCompileForTests(compileVar)
	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		instanceAdapter,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotCompileBytesToInstance {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotCompileBytesToInstance, code)
		return
	}
}

func TestExecute_withCompile_compileDoesNotExistsInFrame_returnsError(t *testing.T) {
	compileVar := "compileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	frame := stacks.NewFrameForTests()
	instruction := compilers.NewCompilerWithCompileForTests(compileVar)
	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{
			string(compile): compiledInstance,
		},
	)

	application := NewApplication(
		instanceAdapter,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchCompileFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchCompileFromFrame, code)
		return
	}
}

func TestExecute_withDecompile_decompileExistsInFrame_adapterSucceeds_Success(t *testing.T) {
	decompileVar := "decompileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				decompileVar,
				stacks.NewAssignableWithInstanceForTests(
					compiledInstance,
				),
			),
		}),
	)

	instruction := compilers.NewCompilerWithDecompileForTests(decompileVar)
	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{
			compiledInstance.Hash().String(): compile,
		},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		instanceAdapter,
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retBytes := retAssignable.Bytes()
	if !bytes.Equal(retBytes, compile) {
		t.Errorf("the returned data is invalid")
		return
	}
}

func TestExecute_withDecompile_decompileExistsInFrame_adapterFails_returnsError(t *testing.T) {
	decompileVar := "decompileVar"
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				decompileVar,
				stacks.NewAssignableWithInstanceForTests(
					compiledInstance,
				),
			),
		}),
	)

	instruction := compilers.NewCompilerWithDecompileForTests(decompileVar)
	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		instanceAdapter,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotDecompileInstanceToBytes {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotDecompileInstanceToBytes, code)
		return
	}
}

func TestExecute_withDecompile_decompileDoesNotExistsInFrame_returnsError(t *testing.T) {
	decompileVar := "decompileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	frame := stacks.NewFrameForTests()
	instruction := compilers.NewCompilerWithDecompileForTests(decompileVar)
	instanceAdapter := mocks.NewInstanceAdapter(
		map[string][]byte{
			compiledInstance.Hash().String(): compile,
		},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		instanceAdapter,
	)

	_, pCode, err := application.Execute(frame, instruction)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	if pCode == nil {
		t.Errorf("the code was expected to be valid, nil returned")
		return
	}

	code := *pCode
	if code != failures.CouldNotFetchDecompileFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchDecompileFromFrame, code)
		return
	}
}
