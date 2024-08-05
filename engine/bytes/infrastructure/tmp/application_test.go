package tmp

import (
	"os"
	"path/filepath"
	"testing"

	infra_bytes "github.com/steve-care-software/webx/engine/bytes/infrastructure/bytes"
)

func TestApplication_Namespaces_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	dbName := "myDatabase"

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	applicationBuilder := NewApplicationBuilder(
		infra_bytes.NewNamespaceAdapter(),
	)

	application, err := applicationBuilder.Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContext, err := application.Begin(dbName, "myNamespace")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	application.Close(*pContext)
}
