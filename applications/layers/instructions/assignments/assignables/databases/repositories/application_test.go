package repositories

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/applications/mocks"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

func TestExecute_withSkeleton_Success(t *testing.T) {
	skeleton := mocks.NewSkeleton()
	frame := stacks.NewFrameForTests()
	instruction := repositories.NewRepositoryWithSkeletonForTests()
	repository := mocks.NewInstanceRepository(
		23,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		repository,
		skeleton,
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

	retSkeleton := retAssignable.Instance().(skeletons.Skeleton)
	if !reflect.DeepEqual(skeleton, retSkeleton) {
		t.Errorf("the returned skeleton is invalid")
		return
	}
}

func TestExecute_withHeight_Success(t *testing.T) {
	height := uint(43)
	frame := stacks.NewFrameForTests()
	instruction := repositories.NewRepositoryWithHeightForTests()
	repository := mocks.NewInstanceRepository(
		height,
		[]hash.Hash{},
		map[string]instances.Instance{},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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
		t.Errorf("the assignable was expected to contain a uint")
		return
	}

	pRetValue := retAssignable.UnsignedInt()
	if *pRetValue != height {
		t.Errorf("the returned height is invalid")
		return
	}
}

func TestExecute_withList_listExistsInFrame_listSucceeds_Success(t *testing.T) {
	listVar := "myList"
	query := mocks.NewQuery()

	firstHash, err := hash.NewAdapter().FromBytes([]byte("this is some hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashList := []hash.Hash{
		*firstHash,
	}

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				listVar,
				stacks.NewAssignableWithQueryForTests(
					query,
				),
			),
		}),
	)

	instruction := repositories.NewRepositoryWithListForTests(listVar)
	repository := mocks.NewInstanceRepository(
		uint(43),
		hashList,
		map[string]instances.Instance{},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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

	if !retAssignable.IsHashList() {
		t.Errorf("the assignable was expected to contain an hash list")
		return
	}

	retHashList := retAssignable.HashList()
	if !reflect.DeepEqual(hashList, retHashList) {
		t.Errorf("the returned hash list is invalid")
		return
	}
}

func TestExecute_withList_listExistsInFrame_listFails_returnsError(t *testing.T) {
	listVar := "myList"
	query := mocks.NewQuery()

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				listVar,
				stacks.NewAssignableWithQueryForTests(
					query,
				),
			),
		}),
	)

	instruction := repositories.NewRepositoryWithListForTests(listVar)
	repository := mocks.NewInstanceRepository(
		uint(43),
		nil,
		map[string]instances.Instance{},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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
	if code != failures.CouldNotListInstancesFromDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotListInstancesFromDatabase, code)
		return
	}
}

func TestExecute_withList_listDoesNotExistsInFrame_returnsError(t *testing.T) {
	listVar := "myList"
	firstHash, err := hash.NewAdapter().FromBytes([]byte("this is some hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashList := []hash.Hash{
		*firstHash,
	}

	frame := stacks.NewFrameForTests()
	instruction := repositories.NewRepositoryWithListForTests(listVar)
	repository := mocks.NewInstanceRepository(
		uint(43),
		hashList,
		map[string]instances.Instance{},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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
	if code != failures.CouldNotFetchListQueryFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchListQueryFromFrame, code)
		return
	}
}

func TestExecute_withRetrieve_retrieveExistsInFrame_retrieveSucceeds_Success(t *testing.T) {
	retrieveVar := "myRetrieve"
	query := mocks.NewQuery()
	instance := repositories.NewRepositoryWithRetrieveForTests(retrieveVar)

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				retrieveVar,
				stacks.NewAssignableWithQueryForTests(
					query,
				),
			),
		}),
	)

	instruction := repositories.NewRepositoryWithRetrieveForTests(retrieveVar)
	repository := mocks.NewInstanceRepository(
		uint(43),
		nil,
		map[string]instances.Instance{
			query.Hash().String(): instance,
		},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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
	if !reflect.DeepEqual(instance, retInstance) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestExecute_withRetrieve_retrieveExistsInFrame_retrieveFails_returnsError(t *testing.T) {
	retrieveVar := "myRetrieve"
	query := mocks.NewQuery()
	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				retrieveVar,
				stacks.NewAssignableWithQueryForTests(
					query,
				),
			),
		}),
	)

	instruction := repositories.NewRepositoryWithRetrieveForTests(retrieveVar)
	repository := mocks.NewInstanceRepository(
		uint(43),
		nil,
		map[string]instances.Instance{},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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
	if code != failures.CouldNotRetrieveInstanceFromDatabase {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotRetrieveInstanceFromDatabase, code)
		return
	}
}

func TestExecute_withRetrieve_retrieveDoesNotExistsInFrame_returnsError(t *testing.T) {
	retrieveVar := "myRetrieve"
	query := mocks.NewQuery()
	instance := repositories.NewRepositoryWithRetrieveForTests(retrieveVar)

	frame := stacks.NewFrameForTests()
	instruction := repositories.NewRepositoryWithRetrieveForTests(retrieveVar)
	repository := mocks.NewInstanceRepository(
		uint(43),
		nil,
		map[string]instances.Instance{
			query.Hash().String(): instance,
		},
	)

	application := NewApplication(
		repository,
		mocks.NewSkeleton(),
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
	if code != failures.CouldNotFetchRetrieveQueryFromFrame {
		t.Errorf("the code was expected to be %d, %d returned", failures.CouldNotFetchRetrieveQueryFromFrame, code)
		return
	}
}
