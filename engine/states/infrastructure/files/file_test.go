package files

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestTransactManually_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	innerPath := []string{
		"inner_dir",
	}

	path := []string{
		"path_dir",
		"myFile.data",
	}

	expectedBytes := []byte("this is some data")

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	repository, err := NewRepositoryBuilder(innerPath).Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	service, err := NewServiceBuilder(innerPath).Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if repository.Exists(path) {
		t.Errorf("the path: %s, was expected to NOT exists", filepath.Join(path...))
		return
	}

	err = service.Save(path, expectedBytes)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	err = service.Init(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Save(path, expectedBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !repository.Exists(path) {
		t.Errorf("the path: %s, was expected to exists", filepath.Join(path...))
		return
	}

	retBytes, err := repository.Retrieve(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedBytes, retBytes) {
		t.Errorf("the returned bytes are invalid")
		return
	}
}

func TestTransact_Success(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	innerPath := []string{
		"inner_dir",
	}

	path := []string{
		"path_dir",
		"myFile.data",
	}

	expectedBytes := []byte("this is some data")

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	repository, err := NewRepositoryBuilder(innerPath).Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	service, err := NewServiceBuilder(innerPath).Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if repository.Exists(path) {
		t.Errorf("the path: %s, was expected to NOT exists", filepath.Join(path...))
		return
	}

	err = service.Transact(path, expectedBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !repository.Exists(path) {
		t.Errorf("the path: %s, was expected to exists", filepath.Join(path...))
		return
	}

	retBytes, err := repository.Retrieve(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedBytes, retBytes) {
		t.Errorf("the returned bytes are invalid")
		return
	}

	err = service.Transact(path, expectedBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytesAgain, err := repository.Retrieve(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedBytes, retBytesAgain) {
		t.Errorf("the returned bytes are invalid")
		return
	}
}

func TestTransact_tryLockTwice_returnsError(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	innerPath := []string{
		"inner_dir",
	}

	path := []string{
		"path_dir",
		"myFile.data",
	}

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	service, err := NewServiceBuilder(innerPath).Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Init(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Lock(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Lock(path)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTransact_unlockAnNonLockFile_returnsError(t *testing.T) {
	basePath := []string{
		"test_files",
	}

	innerPath := []string{
		"inner_dir",
	}

	path := []string{
		"path_dir",
		"myFile.data",
	}

	defer func() {
		os.RemoveAll(filepath.Join(basePath...))
	}()

	service, err := NewServiceBuilder(innerPath).Create().WithBasePath(basePath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Init(path)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = service.Unlock(path)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
