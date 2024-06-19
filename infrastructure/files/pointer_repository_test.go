package files

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers"
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
	json_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers"
)

func TestPointerRepository_Retrieve_Success(t *testing.T) {
	repositoryBuilder := NewPointerRepositoryBuilder(
		json_pointers.NewAdapter(),
		pointers.NewBuilder(),
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pointers, err := repository.Retrieve([]string{
		"pointers_layer.json",
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := pointers.List()
	if len(list) != 2 {
		t.Errorf("the pointers file was expected to contain %d pointers, %d returned", 2, len(list))
		return
	}
}

func TestPointerRepository_Retrieve_withInvalidPath_returnsError(t *testing.T) {
	repositoryBuilder := NewPointerRepositoryBuilder(
		json_pointers.NewAdapter(),
		pointers.NewBuilder(),
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = repository.Retrieve([]string{
		"invalid_path.json",
	})

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestPointerRepository_Match_matchesOnePointerOutOfTwo_Success(t *testing.T) {
	repositoryBuilder := NewPointerRepositoryBuilder(
		json_pointers.NewAdapter(),
		pointers.NewBuilder(),
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	history := [][]string{
		[]string{"this", "is", "an", "executed", "path"},
	}

	pointers, err := repository.Match(
		[]string{
			"pointers_layer.json",
		},
		history,
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := pointers.List()
	if len(list) != 1 {
		t.Errorf("the pointers file was expected to contain %d pointers, %d returned", 1, len(list))
		return
	}
}

func TestPointerRepository_Match_matchesZeroPointerOutOfTwo_returnsError(t *testing.T) {
	repositoryBuilder := NewPointerRepositoryBuilder(
		json_pointers.NewAdapter(),
		pointers.NewBuilder(),
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	history := [][]string{
		[]string{"this", "is", "an", "executed", "path"},
	}

	_, err = repository.Match(
		[]string{
			"zero_match_pointers_layer.json",
		},
		history,
	)

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestPointerRepository_Fetch_returnsLayerBytes_Success(t *testing.T) {
	repositoryBuilder := NewPointerRepositoryBuilder(
		json_pointers.NewAdapter(),
		pointers.NewBuilder(),
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	history := [][]string{
		[]string{"this", "is", "an", "executed", "path"},
	}

	bytes, err := repository.Fetch(
		[]string{
			"pointers_layer.json",
		},
		history,
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retLayer, err := json_layers.NewAdapter().ToInstance(bytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retLayer == nil {
		t.Errorf("the layer was NOT expected to be nil")
		return
	}
}
