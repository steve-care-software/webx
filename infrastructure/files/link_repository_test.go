package files

import (
	"testing"

	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links"
)

func TestLinkRepository_Success(t *testing.T) {
	adapter := links.NewAdapter()
	repositoryBuilder := NewLinkRepositoryBuilder(
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

	link, err := repository.Retrieve([]string{
		"link.json",
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if link == nil {
		t.Errorf("the link was expected to be valid")
		return
	}
}
