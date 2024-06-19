package files

import (
	"testing"

	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
)

func TestLayerRepository_Success(t *testing.T) {
	adapter := layers.NewAdapter()
	repositoryBuilder := NewLayerRepositoryBuilder(
		adapter,
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	layer, err := repository.Retrieve([]string{
		"layers",
		"first.json",
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if layer == nil {
		t.Errorf("the layer was expected to be valid")
		return
	}
}
