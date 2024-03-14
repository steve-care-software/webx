package libraries

import (
	"reflect"
	"testing"

	application_compilers "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/compilers"
	application_databases "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/databases"
	application_repositories "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/databases/repositories"
	application_services "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_WithCompiler_Success(t *testing.T) {
	compileVar := "compileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	context := uint(45)
	skeleton := mocks.NewSkeleton()

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

	instruction := libraries.NewLibraryWithCompilerForTests(
		compilers.NewCompilerWithCompileForTests(compileVar),
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{
					string(compile): compiledInstance,
				},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

func TestExecute_WithDatabase_Success(t *testing.T) {
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	context := uint(45)
	skeleton := mocks.NewSkeleton()

	frame := stacks.NewFrameForTests()
	instruction := libraries.NewLibraryWithDatabaseForTests(
		databases.NewDatabaseWithRepositoryForTests(
			repositories.NewRepositoryWithSkeletonForTests(),
		),
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{
					string(compile): compiledInstance,
				},
			),
		),
		application_databases.NewApplication(
			application_repositories.NewApplication(
				mocks.NewInstanceRepository(
					23,
					[]hash.Hash{},
					map[string]instances.Instance{},
				),
				skeleton,
			),
			application_services.NewApplication(
				mocks.NewInstanceService(
					&context,
					false,
				),
			),
		),
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

	if !retAssignable.IsSkeleton() {
		t.Errorf("the assignable was expected to contain a skeleton")
		return
	}

	retSkeleton := retAssignable.Skeleton()
	if !reflect.DeepEqual(skeleton, retSkeleton) {
		t.Errorf("the returned skeleton is invalid")
		return
	}
}
