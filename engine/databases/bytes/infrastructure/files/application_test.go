package files

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
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

	// insert the entry:
	firstDelimiter, err := application.Insert(*pContext, firstData)
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

	// insert the second entry:
	secondDelimiter, err := application.Insert(*pContext, secondData)
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

	// retrieve the deleted state indexes:
	deletedIndexes, err := application.DeletedStateIndexes(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(deletedIndexes) != 1 {
		t.Errorf("%d deleted states were expected, %d returned", len(deletedIndexes), 1)
		return
	}

	if deletedIndexes[0] != 2 {
		t.Errorf("the state %d was expected to be deleted, %d was actually deleted", 2, deletedIndexes[0])
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

	// verify the amount of states:
	pAmountStates, err := application.StatesAmount(*pContext)
	if err != nil {
		t.Errorf("the error was expected to contain an error, nil returned")
		return
	}

	if *pAmountStates != 3 {
		t.Errorf("%d states were expected, %d returned", 3, *pAmountStates)
		return
	}

}
