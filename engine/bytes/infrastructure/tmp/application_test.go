package tmp

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces/updates"
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

	// begin:
	pContext, err := application.Begin(dbName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// insert namespace:
	firstName := "first"
	firstDescription := "This is a description"
	err = application.InsertNamespace(*pContext, firstName, firstDescription)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// set the namespace:
	err = application.SetNamespace(*pContext, firstName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// update the namespace:
	firstUpdatedName := "firstUpdatedName"
	updatedNamespace, err := updates.NewBuilder().Create().WithName(firstUpdatedName).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.UpdateNamespace(*pContext, firstName, updatedNamespace)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// delete the namespace:
	err = application.DeleteNamespace(*pContext, firstUpdatedName)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// commit:
	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
