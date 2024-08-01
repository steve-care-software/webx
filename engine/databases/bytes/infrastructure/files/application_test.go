package files

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

func TestSingleTransaction_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	dbName := "mydatabase.db"
	application, err := NewApplicationBuilder().Create().WithBasePath(basePath).Now()
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

	defer application.Close(*pContext)

	// create an entry:
	firstData := []byte("this is some data")
	firstDelimiter := delimiters.NewDelimiterForTests(0, uint64(len(firstData)))
	firstEntry := entries.NewEntryForTests(
		firstDelimiter,
		firstData,
	)

	// insert the entry:
	err = application.Insert(*pContext, firstEntry)
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

	// create a second entry entry:
	secondData := []byte("this is some other data")
	secondDelimiter := delimiters.NewDelimiterForTests(
		firstDelimiter.Index()+firstDelimiter.Length(),
		uint64(len(secondData)),
	)

	secondEntry := entries.NewEntryForTests(
		secondDelimiter,
		secondData,
	)

	// insert the second entry:
	err = application.Insert(*pContext, secondEntry)
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

	// retrieve the content from the first delimiter:
	retData, err := application.Retrieve(*pContext, firstDelimiter)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// the returned data is expected to be equal:
	if !bytes.Equal(firstData, retData) {
		t.Errorf("the returned data is invalid")
		return
	}

	// delete an entry:
	err = application.Delete(*pContext, firstDelimiter)
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

	// retrieve the content from the first delimiter, returns error:
	_, err = application.Retrieve(*pContext, firstDelimiter)
	if err == nil {
		t.Errorf("the error was expected to contain an error, nil returned")
		return
	}

	// retrieve the content from the second delimiter:
	retData, err = application.Retrieve(*pContext, secondDelimiter)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// the returned data is expected to be equal:
	if !bytes.Equal(secondData, retData) {
		t.Errorf("the returned data is invalid")
		return
	}

	// delete the third state:
	err = application.DeleteState(*pContext, 2)
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

	// retrieve the content from the second delimiter, returns error:
	_, err = application.Retrieve(*pContext, secondDelimiter)
	if err == nil {
		t.Errorf("the error was expected to contain an error, nil returned")
		return
	}

	// recover the third state:
	err = application.RecoverState(*pContext, 2)
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

	// retrieve the content from the second delimiter:
	retData, err = application.Retrieve(*pContext, secondDelimiter)
	if err != nil {
		t.Errorf("the error was expected to contain an error, nil returned")
		return
	}

	// the returned data is expected to be equal:
	if !bytes.Equal(secondData, retData) {
		t.Errorf("the returned data is invalid")
		return
	}

}
