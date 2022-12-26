package files

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

func TestExists_thenCreate_thenDelete_Success(t *testing.T) {
	miningValue := []byte("0")[0]
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	application := NewApplication(miningValue, dirPath, dstExtension, bckExtension, readChunkSize)

	name := "my_name"
	exists, err := application.Exists(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the database was expected to NOT exists")
		return
	}

	err = application.New(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	exists, err = application.Exists(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !exists {
		t.Errorf("the database was expected to exists")
		return
	}

	err = application.New(name)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	err = application.Delete(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	exists, err = application.Exists(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the database was expected to NOT exists")
		return
	}
}

func TestCreate_thenOpen_thenConnections_thenWrite_thenRead_Success(t *testing.T) {
	miningValue := []byte("0")[0]
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	hashAdapter := hash.NewAdapter()
	application := NewApplication(miningValue, dirPath, dstExtension, bckExtension, readChunkSize)

	name := "my_name"
	err := application.New(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pContext, err := application.Open(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	data := []byte("this is some data")
	pHash, err := hashAdapter.FromBytes([]byte("first data hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	kind := uint(0)
	err = application.Write(*pContext, *pHash, data, kind)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retData, err := application.ReadByHash(*pContext, *pHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(retData, data) != 0 {
		t.Errorf("the returned data is invalid")
		return
	}

	retContentKeys, err := application.ContentKeysByKind(*pContext, kind)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retContentKeysList := retContentKeys.List()
	if len(retContentKeysList) != 1 {
		t.Errorf("%d contentKeys od kinf (%d) were expected, %d returned", kind, 1, len(retContentKeysList))
		return
	}

	invalidKind := uint(2345234)
	_, err = application.ContentKeysByKind(*pContext, invalidKind)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	retCommits, err := application.Commits(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := retCommits.List()
	if len(list) != 1 {
		t.Errorf("%d commits were expected, %d returned", 1, len(list))
		return
	}

	retCommit, err := application.CommitByHash(*pContext, list[0].Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(retCommit, list[0]) {
		t.Errorf("the returned commit is invalid")
		return
	}

	err = application.Close(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondContext, err := application.Open(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondHash, err := hashAdapter.FromBytes([]byte("second data hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	secondData := []byte("this is some second additional data")
	err = application.Write(*pSecondContext, *pSecondHash, secondData, kind)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit(*pSecondContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retSecondData, err := application.ReadByHash(*pSecondContext, *pSecondHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Close(*pSecondContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(retSecondData, secondData) != 0 {
		t.Errorf("the returned data is invalid")
		return
	}

	pFourthContext, err := application.Open(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFirstData, err := application.ReadByHash(*pFourthContext, *pHash)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(retFirstData, data) != 0 {
		t.Errorf("the returned data is invalid")
		return
	}
}

func TestConnections_isEmpty_returnsError(t *testing.T) {
	miningValue := []byte("0")[0]
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	readChunkSize := uint(1000000)
	defer func() {
		os.RemoveAll(dirPath)
	}()

	application := NewApplication(miningValue, dirPath, dstExtension, bckExtension, readChunkSize)

	_, err := application.Connections()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
