package files

import (
	"bytes"
	"os"
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

	kind := uint(0)
	data := []byte("this is some data")
	pHash, err := hashAdapter.FromBytes([]byte("first data hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

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

	err = application.Close(*pContext)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if bytes.Compare(retData, data) != 0 {
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
