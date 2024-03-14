package databases

import (
	"reflect"
	"testing"

	application_repositories "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases/repositories"
	application_services "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_WithRepository_Success(t *testing.T) {
	context := uint(45)
	skeleton := mocks.NewSkeleton()
	frame := stacks.NewFrameForTests()

	instruction := databases.NewDatabaseWithRepositoryForTests(
		repositories.NewRepositoryWithSkeletonForTests(),
	)

	application := NewApplication(
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

func TestExecute_WithService_Success(t *testing.T) {
	context := uint(45)
	frame := stacks.NewFrameForTests()
	instruction := databases.NewDatabaseWithServiceForTests(
		services.NewServiceWithBeginForTests(),
	)

	application := NewApplication(
		application_repositories.NewApplication(
			mocks.NewInstanceRepository(
				23,
				[]hash.Hash{},
				map[string]instances.Instance{},
			),
			mocks.NewSkeleton(),
		),
		application_services.NewApplication(
			mocks.NewInstanceService(
				&context,
				false,
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

	if !retAssignable.IsUnsignedInt() {
		t.Errorf("the assignable was expected to contain an uint")
		return
	}

	pRetValue := retAssignable.UnsignedInt()
	if !reflect.DeepEqual(context, *pRetValue) {
		t.Errorf("the returned context is invalid")
		return
	}
}
