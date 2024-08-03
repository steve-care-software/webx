package applications

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	bytesdb_infra_files "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/files"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	infra_bytes "github.com/steve-care-software/webx/engine/databases/hashes/infrastructure/bytes"
)

func TestApplication_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	hashAdapter := hash.NewAdapter()
	bytesApp, err := bytesdb_infra_files.NewApplicationBuilder().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	application, err := NewBuilder(
		infra_bytes.NewPointerAdapter(),
	).
		Create().
		WithBytes(bytesApp).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	name := "someDatabase"
	pContext, err := application.Begin(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstData := []byte("this is some data")
	pFirstHash, err := hashAdapter.FromBytes([]byte("this is the first key"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = application.Retrieve(*pContext, *pFirstHash)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	err = application.Insert(*pContext, *pFirstHash, firstData)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFirstData, err := application.Retrieve(*pContext, *pFirstHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(firstData, retFirstData) {
		t.Errorf("the returned data is invalid")
		return
	}

	secondData := []byte("this is some second data")
	pSecondHash, err := hashAdapter.FromBytes([]byte("this is the second key"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Insert(*pContext, *pSecondHash, secondData)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Delete(*pContext, *pFirstHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = application.Retrieve(*pContext, *pFirstHash)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	retSecondData, err := application.Retrieve(*pContext, *pSecondHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(secondData, retSecondData) {
		t.Errorf("the returned data is invalid")
		return
	}

	err = application.Close(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
