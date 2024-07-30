package files

import (
	"testing"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"
)

func TestSingleTransaction_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

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

	// create an entry:
	data := []byte("this is some data")
	entry := entries.NewEntryForTests(
		"myContainer",
		delimiters.NewDelimiterForTests(0, uint64(len(data))),
		data,
	)

	// insert the entry:
	err = application.Insert(*pContext, entry)
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
