package files

import (
	"bytes"
	"math/rand/v2"
	"os"
	"path/filepath"
	"testing"

	applications_loaders "github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/namespaces/versions/workspaces/branches/states/pointers"
	applications_writers "github.com/steve-care-software/webx/engine/cursors/applications/sessions/writers/namespaces/versions/workspaces/branches/states/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

func TestApplication_withPointers_Success(t *testing.T) {
	originalPath := []string{
		"test_files",
		"original.db",
	}

	destinationPath := []string{
		"test_files",
		"destination.db",
	}

	originalPathStr := filepath.Join(originalPath...)
	defer func() {
		os.Remove(originalPathStr)
		os.Remove(filepath.Join(destinationPath...))
	}()

	// create the first data:
	firstData := []byte("this is some original data")

	// generate some data:
	amount := int((readChunkSize * 2) + 23)
	originalBytes := []byte{}
	for i := 0; i < amount; i++ {
		value := uint8(rand.IntN(255))
		originalBytes = append(originalBytes, value)
	}

	// add some empty data to the original bytes:
	emptyData := make([]byte, len(firstData))
	updatedOriginalBytes := append(originalBytes, emptyData...)

	// add some deletedData data:
	deletedData := []byte{}
	amount = int((readChunkSize * 2) + 23)
	for i := 0; i < amount; i++ {
		value := uint8(rand.IntN(255))
		deletedData = append(deletedData, value)
	}

	// write the to the original file:
	err := os.WriteFile(originalPathStr, append(updatedOriginalBytes, deletedData...), 0777)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// create the new pointers:
	firstIndex := uint64(len(originalBytes))
	firstLength := uint64(len(firstData))
	pointers := pointers.NewPointersForTests([]pointers.Pointer{
		pointers.NewPointerForTests(
			storages.NewStorageForTests(
				delimiters.NewDelimiterForTests(firstIndex, firstLength),
				false,
			),
			[]byte(firstData),
		),
		pointers.NewPointerForTests(
			storages.NewStorageForTests(
				delimiters.NewDelimiterForTests(firstIndex+firstLength, uint64(len(deletedData))),
				true, // deleted
			),
			[]byte(deletedData),
		),
	})

	loaderApp := applications_loaders.NewApplication()
	fileApp, err := NewDatabaseApplicationBuilder().Create().WithOriginal(originalPath).WithDestination(destinationPath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	defer fileApp.CloseThenDeleteOriginal()

	writerApp := applications_writers.NewApplication(
		fileApp,
	)

	secondData := []byte("this is some inserted data")
	retLoadedPointers, err := loaderApp.InsertData(pointers, secondData)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = writerApp.Write(uint64(0), retLoadedPointers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// read the destination file:
	retDestinationBytes, err := os.ReadFile(filepath.Join(destinationPath...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := append(originalBytes, firstData...)
	expected = append(expected, deletedData...)
	expected = append(expected, secondData...)
	if !bytes.Equal(retDestinationBytes, expected) {
		t.Errorf("the returned data is invalid")
		return
	}
}
