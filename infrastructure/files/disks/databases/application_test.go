package databases

import (
	"os"
	"testing"
)

func TestNew_Success(t *testing.T) {
	dirPath := "./test_files"
	defer func() {
		os.RemoveAll(dirPath)
	}()

	application, err := NewBuilder().Create().WithDirPath(dirPath).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	beforeList, err := application.List()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	dbFile := "mydatabase.db"
	err = application.New(dbFile)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	afterList, err := application.List()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(beforeList) != 0 {
		t.Errorf("%d databases were expected at the beginning, %d returned", 0, len(beforeList))
		return
	}

	if len(afterList) != 1 {
		t.Errorf("%d databases were expected at the beginning, %d returned", 1, len(afterList))
		return
	}
}
