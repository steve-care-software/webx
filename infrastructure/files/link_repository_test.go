package files

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/pointers"
	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links"
	json_pointers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers"
)

func TestLinkRepository_Success(t *testing.T) {
	repositoryBuilder := NewLinkRepositoryBuilder(
		NewPointerRepositoryBuilder(
			json_pointers.NewAdapter(),
			pointers.NewBuilder(),
		),
		links.NewAdapter(),
	)

	basePath := []string{
		"test_files",
	}

	repository, err := repositoryBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	link, err := repository.Retrieve(
		[]string{
			"pointers_link.json",
		},
		[][]string{},
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if link == nil {
		t.Errorf("the link was expected to be valid")
		return
	}
}
